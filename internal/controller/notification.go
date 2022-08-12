package controller

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"notification-service/internal/util/iface"
	"strings"

	"notification-service/internal/dto"
	"notification-service/internal/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util"
)

type NotificationV1Controller interface {
	HandleAll(res http.ResponseWriter, req *http.Request)
}

type basicNotificationV1Controller struct {
	templateRepository     repository.TemplateRepository
	notificationRepository repository.NotificationRepository
}

func NewNotificationV1Controller(templateRepository repository.TemplateRepository,
							   notificationRepository repository.NotificationRepository) NotificationV1Controller {
	return &basicNotificationV1Controller{
		templateRepository,
		notificationRepository,
	}
}

func (bnc *basicNotificationV1Controller) HandleAll(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)

	switch req.Method {
		case http.MethodGet: {
			bnc.getBulk(brw, req)
		}
		case http.MethodPost: {
			bnc.send(brw, req)
		}
		default: {
			brw.Status(http.StatusMethodNotAllowed)
		}
	}
}

func (bnc *basicNotificationV1Controller) getBulk(res iface.IResponseWriter, req *http.Request) {
	// GET /notifications
	// GET /notifications?page=24 (size = default = 20)
	// GET /notifications?size=50 (page = default = 1)
	// GET /notifications?appId=aa-bb
	// GET /notifications?templateId=45
	// GET /notifications?startTime=17824254
	// GET /notifications?endTime=17824254

	filter := entity.NotificationFilterFromRequest(req, res)
	if filter == nil {
		return
	}

	notifications := bnc.notificationRepository.GetBulk(filter)
	if notifications == nil {
		res.Status(http.StatusBadRequest).Text("Failed to get anything")
	} else {
		res.Status(http.StatusOK).Json(*notifications)
	}
}

func (bnc *basicNotificationV1Controller) send(res iface.IResponseWriter, req *http.Request) {
	reqObj := dto.SendNotificationRequest{}
	if !util.ConvertFromJson(res, req, &reqObj) {
		return
	}

	templateEntity, status := bnc.templateRepository.Get(*reqObj.TemplateID)
	if status == 1 {
		res.Status(http.StatusNotFound).Text("Something was wrong with the database. Try again")
		return
	} else if status == 2 {
		res.Status(http.StatusNotFound).Text("Template was not found!")
		return
	}

	if templateEntity.ContactType != *reqObj.ContactType {
		res.Status(http.StatusBadRequest).Text("'contactType' should be '" + templateEntity.ContactType + "' in order to use this template")
		return
	}

	for i := 0; i < len(reqObj.Placeholders); i++ {
		placeholder := &(reqObj.Placeholders[i])
		key := "@{" + (*placeholder.Key) + "}"
		templateEntity.Template = strings.ReplaceAll(templateEntity.Template, key, *placeholder.Value)
	}

	unfilledPlaceholders := templateEntity.GetPlaceholders()
	if len(unfilledPlaceholders) > 0 {
		log.Error("Unfilled placeholders: ", unfilledPlaceholders)
		res.Status(http.StatusUnprocessableEntity).Text("Unfilled placeholders: " + unfilledPlaceholders)
		return
	}

	notificationEntity := entity.NotificationEntity {
		TemplateID:           *reqObj.TemplateID,
		AppID:                *reqObj.AppID,
		ContactType:          *reqObj.ContactType,
		ContactInfo:          *reqObj.GetContactInfo(),
		Title:                *reqObj.Title,
		Message:              templateEntity.Template,
	}
	
	status2 := bnc.notificationRepository.Insert(&notificationEntity)
	if status2 {
		res.Status(http.StatusCreated).Text("Notification created successfully!")
	} else {
		res.Status(http.StatusBadRequest).Text("Failed to create notification. Try again!")
	}
}