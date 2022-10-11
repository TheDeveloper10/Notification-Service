package controller

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller/layer"
	"notification-service/internal/dto"
	"notification-service/internal/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"strconv"
	"sync"
)

type NotificationV1Controller interface {
	iface.IController
	CreateNotificationFromBytes(bytes []byte)
}

func NewNotificationV1Controller(templateRepository repository.ITemplateRepository,
								notificationRepository repository.INotificationRepository,
								clientRepository repository.IClientRepository) NotificationV1Controller {
	return &basicNotificationV1Controller{
		templateRepository,
		notificationRepository,
		clientRepository,
	}
}

type basicNotificationV1Controller struct {
	templateRepository     repository.ITemplateRepository
	notificationRepository repository.INotificationRepository
	clientRepository       repository.IClientRepository
}

func (bnc *basicNotificationV1Controller) CreateRoutes(router *rem.Router) {
	router.
		NewRoute("/v1/notifications").
		Get(bnc.getBulk).
		Post(bnc.send)
}

func (bnc *basicNotificationV1Controller) CreateNotificationFromBytes(bytes []byte) {
	reqObj := dto.SendNotificationRequest{}
	if !layer.JSONBytesConverterMiddleware(bytes, &reqObj) {
		return
	}

	res := util.LogDataResponseWriter{}
	bnc.internalSend(&reqObj, &res)
}

func (bnc *basicNotificationV1Controller) getBulk(res rem.IResponse, req rem.IRequest) bool {
	// GET /notifications
	// GET /notifications?page=24 (size = default = 20)
	// GET /notifications?size=50 (page = default = 1)
	// GET /notifications?appId=aa-bb
	// GET /notifications?templateId=45
	// GET /notifications?startTime=17824254
	// GET /notifications?endTime=17824254
	if !layer.AccessTokenMiddleware(bnc.clientRepository, res, req, entity.PermissionReadSentNotifications) {
		return true
	}

	filter := entity.NotificationFilterFromRequest(req, res)
	if filter == nil {
		return true
	}

	notifications := bnc.notificationRepository.GetBulk(filter)
	if notifications == nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError("Failed to get anything"))
	} else if len(*notifications) > 0 {
		res.Status(http.StatusOK).JSON(*notifications)
	} else {
		res.Status(http.StatusOK)
	}

	return true
}

func (bnc *basicNotificationV1Controller) send(res rem.IResponse, req rem.IRequest) bool {
	if !layer.AccessTokenMiddleware(bnc.clientRepository, res, req, entity.PermissionSendNotifications) {
		return true
	}

	reqObj := dto.SendNotificationRequest{}
	if !layer.JSONConverterMiddleware(res, req, &reqObj) {
		return true
	}

	bnc.internalSend(&reqObj, res)
	return true
}

func (bnc *basicNotificationV1Controller) internalSend(reqObj *dto.SendNotificationRequest, res rem.IResponse) {
	templateEntity, status := bnc.templateRepository.Get(reqObj.TemplateID)
	if status == 1 {
		res.Status(http.StatusNotFound).JSON(util.ErrorListFromTextError("Something was wrong with the database. Try again"))
		return
	} else if status == 2 {
		res.Status(http.StatusNotFound).JSON(util.ErrorListFromTextError("Template was not found!"))
		return
	}

	filledResult := bnc.fillPlaceholders(templateEntity, &reqObj.UniversalPlaceholders, res)
	if !filledResult {
		return
	}

	var wg sync.WaitGroup
	errs := make([]dto.SendNotificationErrorData, 5)

	sentNotifications := 0
	failedNotifications := 0
	targetCount := len(reqObj.Targets)
	for currentTarget := 0; currentTarget < targetCount; currentTarget++ {
		target := &(reqObj.Targets[currentTarget])

		targetErrors := target.Validate()
		if targetErrors != nil && targetErrors.ErrorsCount() > 0 {
			errs = append(errs, dto.SendNotificationErrorData{
				TargetId: currentTarget,
				Messages: *targetErrors.GetErrors(),
			})
			continue
		}

		if target.Email != nil && !bnc.sendEmail(currentTarget, templateEntity, reqObj, &errs, &wg) {
			continue
		}

		if target.PhoneNumber != nil && !bnc.sendSMS(currentTarget, templateEntity, reqObj, &errs, &wg) {
			continue
		}

		if target.FCMRegistrationToken != nil && !bnc.sendPush(currentTarget, templateEntity, reqObj, &errs, &wg) {
			continue
		}
	}

	wg.Wait()

	if len(errs) > 0 {
		res.Status(http.StatusBadRequest).JSON(dto.SendNotificationsError{
			Errors: errs,
			SuccessfullySentNotifications: sentNotifications,
			FailedNotifications: failedNotifications,
		})
	} else {
		res.Status(http.StatusCreated).Text(strconv.Itoa(targetCount) + " notification(s) have been sent successfully!")
	}
}

func (bnc *basicNotificationV1Controller) sendEmail(targetId int, template *entity.TemplateEntity, request *dto.SendNotificationRequest, errs *[]dto.SendNotificationErrorData, wg *sync.WaitGroup) bool {
	notification := bnc.toNotificationEntity(targetId, template, request, errs)
	if notification == nil {
		return false
	}
	notification.ContactInfo = *request.Targets[targetId].Email

	go bnc.notificationRepository.SendEmail(notification)
	wg.Add(1)
	return true
}

func (bnc *basicNotificationV1Controller) sendSMS(targetId int, template *entity.TemplateEntity, request *dto.SendNotificationRequest, errs *[]dto.SendNotificationErrorData, wg *sync.WaitGroup) bool {
	notification := bnc.toNotificationEntity(targetId, template, request, errs)
	if notification == nil {
		return false
	}
	notification.ContactInfo = *request.Targets[targetId].PhoneNumber

	go bnc.notificationRepository.SendSMS(notification)
	wg.Add(1)
	return true
}

func (bnc *basicNotificationV1Controller) sendPush(targetId int, template *entity.TemplateEntity, request *dto.SendNotificationRequest, errs *[]dto.SendNotificationErrorData, wg *sync.WaitGroup) bool {
	notification := bnc.toNotificationEntity(targetId, template, request, errs)
	if notification == nil {
		return false
	}
	notification.ContactInfo = *request.Targets[targetId].FCMRegistrationToken

	go bnc.notificationRepository.SendPush(notification)
	wg.Add(1)
	return true
}

func (bnc *basicNotificationV1Controller) toNotificationEntity(targetId int, template *entity.TemplateEntity, request *dto.SendNotificationRequest, errs *[]dto.SendNotificationErrorData) *entity.NotificationEntity {
	replaced, err := dto.FillPlaceholders(*template.Body.Email, &request.Targets[targetId].Placeholders)
	if err != nil {
		*errs = append(*errs, dto.SendNotificationErrorData{
			TargetId: targetId,
			Messages: []string { err.Error() },
		})
		return nil
	}

	unfilledPlaceholders := dto.GetPlaceholders(replaced)

	if len(unfilledPlaceholders) > 0 {
		*errs = append(*errs, dto.SendNotificationErrorData{
			TargetId: targetId,
			Messages: []string { "Unfilled placeholders: " + unfilledPlaceholders },
		})

		return nil
	}

	return &entity.NotificationEntity{
		TemplateID:  template.Id,
		AppID:       request.AppID,
		Title:       request.Title,
		Message:     *replaced,
	}
}

func (bnc *basicNotificationV1Controller) fillPlaceholders(template *entity.TemplateEntity, placeholders *[]dto.TemplatePlaceholder, res rem.IResponse) bool {
	if(template.Body.Email != nil) {
		edited, err := dto.FillPlaceholders(*template.Body.Email, placeholders)
		if err != nil {
			res.Status(http.StatusUnprocessableEntity).JSON(util.ErrorListFromTextError(err.Error()))
			return false
		}
		template.Body.Email = edited
	}

	if(template.Body.SMS != nil) {
		edited, err := dto.FillPlaceholders(*template.Body.SMS, placeholders)
		if err != nil {
			res.Status(http.StatusUnprocessableEntity).JSON(util.ErrorListFromTextError(err.Error()))
			return false
		}
		template.Body.SMS = edited
	}

	if(template.Body.Push != nil) {
		edited, err := dto.FillPlaceholders(*template.Body.Push, placeholders)
		if err != nil {
			res.Status(http.StatusUnprocessableEntity).JSON(util.ErrorListFromTextError(err.Error()))
			return false
		}
		template.Body.Push = edited
	}

	return true
}

func (bnc *basicNotificationV1Controller) processNotificationEntity(
		outsourceNotification func(*entity.NotificationEntity) bool,
		notificationEntity *entity.NotificationEntity,
		processId int,
		failures **string,
		failedCount *int,
		wg *sync.WaitGroup) {
	defer wg.Done()
	if outsourceNotification(notificationEntity) &&
		bnc.notificationRepository.Insert(notificationEntity) {
		return
	}

	(*failedCount)++
	if *failures != nil {
		newFailures := (**failures) + ", " + strconv.Itoa(processId)
		*failures = &newFailures
	} else {
		newFailures := strconv.Itoa(processId)
		*failures = &newFailures
	}
}
