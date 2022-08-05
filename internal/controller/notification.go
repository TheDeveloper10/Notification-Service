package controller

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"notification-service.com/packages/internal/dto"
	"notification-service.com/packages/internal/entity"
	"notification-service.com/packages/internal/repository"
	"notification-service.com/packages/internal/util"
)

type basicNotificationController struct {
	templateRepository     repository.TemplateRepository
	notificationRepository repository.NotificationRepository
}

func NewNotificationController(templateRepository repository.TemplateRepository, 
							   notificationRepository repository.NotificationRepository) Controller {
	return &basicNotificationController{
		templateRepository,
		notificationRepository,
	}
}

func (bnc *basicNotificationController) Handle(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)

	switch req.Method {
		case http.MethodPost: {
			bnc.send(brw, req)
		}
	}
}

func (bnc *basicNotificationController) send(res util.IResponseWriter, req *http.Request) {
	reqObj := dto.SendNotificationRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	record, status := bnc.templateRepository.Get(*reqObj.TemplateID)
	if status == 1 {
		res.Status(http.StatusNotFound).Text("Something was wrong with the database. Try again!")
		return
	} else if status == 2 {
		res.Status(http.StatusNotFound).Text("Template was not found!")
		return
	}

	for i := 0; i < len(reqObj.Placeholders); i++ {
		placeholder := &(reqObj.Placeholders[i])
		key := "@{" + (*placeholder.Key) + "}"
		record.Template = strings.ReplaceAll(record.Template, key, *placeholder.Value)
	}

	unfilledPlaceholders := record.GetPlaceholders()
	if len(unfilledPlaceholders) > 0 {
		log.Error("Unfilled placeholders: ", unfilledPlaceholders)
		res.Status(http.StatusUnprocessableEntity).Text("Unfilled placeholders: " + unfilledPlaceholders)
		return
	}

	notificationEntity := entity.NotificationEntity {
		TemplateID: *reqObj.TemplateID,
		UserID: *reqObj.UserID,
		AppID: *reqObj.AppID,
		ContactType: *reqObj.ContactType,
		ContactInfo: *reqObj.ContactInfo,
		Title: *reqObj.Title,
		Message: record.Template,
	}
	
	status2 := bnc.notificationRepository.Insert(&notificationEntity)
	if status2 {
		res.Status(http.StatusCreated).Text("Notification created successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to create notification. Try again!")
	}
}