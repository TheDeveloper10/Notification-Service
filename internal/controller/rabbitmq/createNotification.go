package rabbitmq

import (
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

func NewCreateNotificationV1Controller(templateRepository repository.ITemplateRepository) iface.IRabbitMQController {
	return &basicCreateNotificationV1Controller{
		templateRepository,
	}
}

type basicCreateNotificationV1Controller struct {
	templateRepository     repository.ITemplateRepository
}

func (bcnc *basicCreateNotificationV1Controller) QueueName() string {
	return util.Config.RabbitMQ.TemplatesQueueName
}

func (bcnc *basicCreateNotificationV1Controller) QueueCapacity() int {
	return util.Config.RabbitMQ.TemplatesQueueMax
}

func (bcnc *basicCreateNotificationV1Controller) Handle(bytes []byte) bool {
//	reqObj := dto.SendNotificationRequest{}
//	if !layer.JSONBytesConverterMiddleware(bytes, &reqObj) {
//		return true
//	}
//
//	templateEntity, status := bcnc.templateRepository.Get(reqObj.TemplateID)
//	if status == code.StatusError {
//		return false
//	} else if status == code.StatusNotFound {
//		return true
//	}
//
//	filledResult := bnc.fillPlaceholders(templateEntity, &reqObj.UniversalPlaceholders, res)
//	if !filledResult {
//		return
//	}
//
//	var wg sync.WaitGroup
//	errs := make([]dto.SendNotificationErrorData, 0)
//
//	successCount := 0
//	failedCount := 0
//	targetCount := len(reqObj.Targets)
//	for currentTarget := 0; currentTarget < targetCount; currentTarget++ {
//		target := &(reqObj.Targets[currentTarget])
//
//		targetErrors := target.Validate()
//		targetErrors.Merge(target.ValidateAgainstTemplate(templateEntity))
//		if targetErrors != nil && targetErrors.ErrorsCount() > 0 {
//			errs = append(errs, dto.SendNotificationErrorData{
//				TargetId: currentTarget,
//				Messages: *targetErrors.GetErrors(),
//			})
//			continue
//		}
//
//
//		if target.Email != nil {
//			notification := bnc.toNotificationEntity(currentTarget, templateEntity, templateEntity.Body.Email, reqObj, &errs)
//			if notification != nil {
//				go bnc.sendNotification(*notification, target.Email, bnc.notificationRepository.SendEmail, &successCount, &failedCount, &wg)
//			}
//		}
//
//
//		if target.PhoneNumber != nil {
//			notification := bnc.toNotificationEntity(currentTarget, templateEntity, templateEntity.Body.SMS, reqObj, &errs)
//			if notification != nil {
//				go bnc.sendNotification(*notification, target.PhoneNumber, bnc.notificationRepository.SendSMS, &successCount, &failedCount, &wg)
//			}
//		}
//
//
//		if target.FCMRegistrationToken != nil {
//			notification := bnc.toNotificationEntity(currentTarget, templateEntity, templateEntity.Body.Push, reqObj, &errs)
//			if notification != nil {
//				go bnc.sendNotification(*notification, target.FCMRegistrationToken, bnc.notificationRepository.SendPush, &successCount, &failedCount, &wg)
//			}
//		}
//	}
//
//	wg.Wait()
//
//	if failedCount > 0 || len(errs) > 0 {
//		res.Status(http.StatusBadRequest).JSON(dto.SendNotificationsError{
//			Errors:                        errs,
//			SuccessfullySentNotifications: successCount,
//			FailedNotifications:           failedCount,
//		})
//	} else {
//		res.Status(http.StatusCreated).Text(strconv.Itoa(targetCount) + " notification(s) have been sent successfully!")
//	}
//
//	return true
//}
//
//
//func (bcnc *basicCreateNotificationV1Controller) fillPlaceholders(
//											template *entity.TemplateEntity,
//											placeholders *[]dto.TemplatePlaceholder) bool {
//	if(template.Body.Email != nil) {
//		edited, err := dto.FillPlaceholders(*template.Body.Email, placeholders)
//		if util.ManageError(err) err != nil {
//			res.Status(http.StatusUnprocessableEntity).JSON(util.ErrorListFromTextError(err.Error()))
//			return false
//		}
//		template.Body.Email = edited
//	}
//
//	if(template.Body.SMS != nil) {
//		edited, err := dto.FillPlaceholders(*template.Body.SMS, placeholders)
//		if err != nil {
//			res.Status(http.StatusUnprocessableEntity).JSON(util.ErrorListFromTextError(err.Error()))
//			return false
//		}
//		template.Body.SMS = edited
//	}
//
//	if(template.Body.Push != nil) {
//		edited, err := dto.FillPlaceholders(*template.Body.Push, placeholders)
//		if err != nil {
//			res.Status(http.StatusUnprocessableEntity).JSON(util.ErrorListFromTextError(err.Error()))
//			return false
//		}
//		template.Body.Push = edited
//	}

	return true
}