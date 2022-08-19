package controller

import (
	"net/http"
	"notification-service/internal/dto"
	"notification-service/internal/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"strconv"
	"sync"
)

type NotificationV1Controller interface {
	HandleAll(http.ResponseWriter, *http.Request)
	CreateNotificationFromBytes([]byte) bool
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




func (bnc *basicNotificationV1Controller) CreateNotificationFromBytes(bytes []byte) bool {
	reqObj := dto.SendNotificationRequest{}
	if !util.ConvertFromJsonBytes(bytes, &reqObj) {
		return false
	}

	res := util.StatusOnlyResponseWriter{}
	bnc.internalSend(&reqObj, &res)
	return res.StatusCode != nil && (*res.StatusCode) == 200
}




func (bnc *basicNotificationV1Controller) HandleAll(res http.ResponseWriter, req *http.Request) {
	brw := util.WrapResponseWriter(&res)

	switch req.Method {
	case http.MethodGet:
		bnc.getBulk(brw, req)
	case http.MethodPost:
		bnc.send(brw, req)
	default:
		brw.Status(http.StatusMethodNotAllowed)
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
		res.Status(http.StatusBadRequest).TextError("Failed to get anything")
	} else if len(*notifications) > 0 {
		res.Status(http.StatusOK).Json(*notifications)
	} else {
		res.Status(http.StatusOK)
	}
}

func (bnc *basicNotificationV1Controller) send(res iface.IResponseWriter, req *http.Request) {
	reqObj := dto.SendNotificationRequest{}
	if !util.ConvertFromJsonRequest(res, req, &reqObj) {
		return
	}

	bnc.internalSend(&reqObj, res)
}

func (bnc *basicNotificationV1Controller) internalSend(reqObj *dto.SendNotificationRequest, res iface.IResponseWriter) {
	templateEntity, status := bnc.templateRepository.Get(*reqObj.TemplateID)
	if status == 1 {
		res.Status(http.StatusNotFound).TextError("Something was wrong with the database. Try again")
		return
	} else if status == 2 {
		res.Status(http.StatusNotFound).TextError("Template was not found!")
		return
	}

	if templateEntity.ContactType != *reqObj.ContactType {
		res.Status(http.StatusUnprocessableEntity).TextError("'contactType' should be '" + templateEntity.ContactType + "' in order to use this template")
		return
	}

	universalText, err := dto.FillPlaceholders(templateEntity.Template, &reqObj.UniversalPlaceholders)
	if err != nil {
		res.Status(http.StatusUnprocessableEntity).Error(err)
		return
	}
	universallyUnfilledPlaceholders := dto.GetPlaceholders(&templateEntity.Template)
	needToCheckPlaceholders := len(universallyUnfilledPlaceholders) > 0

	var outsourceNotification func(*entity.NotificationEntity) bool
	switch templateEntity.ContactType {
	case entity.ContactTypeEmail:
		outsourceNotification = bnc.notificationRepository.SendEmail
	case entity.ContactTypePush:
		outsourceNotification = bnc.notificationRepository.SendPush
	case entity.ContactTypeSMS:
		outsourceNotification = bnc.notificationRepository.SendSMS
	}

	var wg sync.WaitGroup
	var failures *string = nil
	failedCount := 0
	additionalError := ""

	targetCount := len(reqObj.Targets)
	currentTarget := 0
	for ; currentTarget < targetCount; currentTarget++ {
		target := &(reqObj.Targets[currentTarget])

		err := target.Validate(&templateEntity.ContactType)
		if err != nil {
			additionalError = err.Error()
			break
		}

		specificText, err := dto.FillPlaceholders(*universalText, &target.Placeholders)
		if err != nil {
			additionalError = err.Error()
			break
		}
		if needToCheckPlaceholders {
			unfilledPlaceholders := dto.GetPlaceholders(specificText)

			if len(unfilledPlaceholders) > 0 {
				additionalError = "Unfilled placeholders: " + unfilledPlaceholders
				break
			}
		}

		notificationEntity := entity.NotificationEntity{
			TemplateID:  *reqObj.TemplateID,
			AppID:       *reqObj.AppID,
			Title:       *reqObj.Title,
			ContactInfo: *target.GetContactInfo(),
			Message:     *specificText,
		}

		wg.Add(1)
		go bnc.processNotificationEntity(
			outsourceNotification,
			&notificationEntity,
			currentTarget,
			&failures,
			&failedCount,
			&wg,
		)
	}

	wg.Wait()

	if additionalError != "" || failures != nil {
		err1 := ""
		if failures != nil {
			err1 = "Failed to send the following notifications: " + (*failures)
		}
		res.Status(http.StatusBadRequest).Json(dto.SentNotificationsError{
			SentNotifications: currentTarget - failedCount,
			Error1:            err1,
			Error2:            additionalError,
		})
	} else {
		res.Status(http.StatusCreated).Text(strconv.Itoa(targetCount) + " notification(s) have been sent successfully!")
	}
}

func (bnc *basicNotificationV1Controller) processNotificationEntity(outsourceNotification func(*entity.NotificationEntity)bool,
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