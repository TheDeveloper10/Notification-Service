package rabbitmqctrl

import (
	"notification-service/internal/controller/common"
	"notification-service/internal/controller/layer"
	"notification-service/internal/data/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/code"
	"notification-service/internal/util/iface"
	"sync"
)

func NewCreateNotificationV1Controller(templateRepository repository.ITemplateRepository,
									notificationRepository repository.INotificationRepository) iface.IRabbitMQController {
	return &basicCreateNotificationV1Controller{
		templateRepository,
		notificationRepository,
	}
}

type basicCreateNotificationV1Controller struct {
	templateRepository     repository.ITemplateRepository
	notificationRepository repository.INotificationRepository
}

func (bcnc *basicCreateNotificationV1Controller) QueueName() string {
	return util.Config.RabbitMQ.NotificationsQueueName
}

func (bcnc *basicCreateNotificationV1Controller) QueueCapacity() int {
	return util.Config.RabbitMQ.NotificationsQueueMax
}

func (bcnc *basicCreateNotificationV1Controller) Handle(bytes []byte) (any, bool) {
	reqObj := dto.SendNotificationRequest{}
	if !layer.JSONBytesConverterMiddleware(bytes, &reqObj) {
		return util.ErrorListFromTextError("Failed to parse JSON"), true
	} else {
		validationErrs := reqObj.Validate()
		if validationErrs.ErrorsCount() > 0 {
			return validationErrs, true
		}
	}

	templateEntity, status := bcnc.templateRepository.Get(reqObj.TemplateID)
	if status == code.StatusError {
		return nil, false
	} else if status == code.StatusNotFound {
		return util.ErrorListFromTextError("Template not found"), true
	}

	err := common.FillPlaceholdersOnTemplate(templateEntity, &reqObj.UniversalPlaceholders)
	if err != nil {
		return util.ErrorListFromTextError(err.Error()), true
	}

	var wg sync.WaitGroup
	errs := make([]dto.SendNotificationErrorData, 0)

	successCount := 0
	failedCount := 0
	targetCount := len(reqObj.Targets)
	for currentTarget := 0; currentTarget < targetCount; currentTarget++ {
		target := &(reqObj.Targets[currentTarget])

		targetErrors := target.Validate()
		targetErrors.Merge(target.ValidateAgainstTemplate(templateEntity))
		if targetErrors != nil && targetErrors.ErrorsCount() > 0 {
			errs = append(errs, dto.SendNotificationErrorData{
				TargetId: currentTarget,
				Messages: *targetErrors.GetErrors(),
			})
			continue
		}


		if target.Email != nil {
			notification := common.ToNotificationEntity(currentTarget, templateEntity, templateEntity.Body.Email, &reqObj, &errs)
			if notification != nil {
				go common.SendNotification(*notification, target.Email, bcnc.notificationRepository.SendEmail, &successCount, &failedCount, bcnc.notificationRepository, &wg)
			}
		}


		if target.PhoneNumber != nil {
			notification := common.ToNotificationEntity(currentTarget, templateEntity, templateEntity.Body.SMS, &reqObj, &errs)
			if notification != nil {
				go common.SendNotification(*notification, target.PhoneNumber, bcnc.notificationRepository.SendSMS, &successCount, &failedCount, bcnc.notificationRepository, &wg)
			}
		}


		if target.FCMRegistrationToken != nil {
			notification := common.ToNotificationEntity(currentTarget, templateEntity, templateEntity.Body.Push, &reqObj, &errs)
			if notification != nil {
				go common.SendNotification(*notification, target.FCMRegistrationToken, bcnc.notificationRepository.SendPush, &successCount, &failedCount, bcnc.notificationRepository, &wg)
			}
		}
	}

	wg.Wait()

	if failedCount > 0 || len(errs) > 0 {
		return &dto.SendNotificationsError{
			Errors:                        errs,
			SuccessfullySentNotifications: successCount,
			FailedNotifications:           failedCount,
		}, true
	}

	return nil, true
}
