package rabbitmq

import (
	"github.com/sirupsen/logrus"
	"notification-service/internal/controller/layer"
	"notification-service/internal/data/dto"
	"notification-service/internal/data/entity"
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
	return util.Config.RabbitMQ.TemplatesQueueName
}

func (bcnc *basicCreateNotificationV1Controller) QueueCapacity() int {
	return util.Config.RabbitMQ.TemplatesQueueMax
}

func (bcnc *basicCreateNotificationV1Controller) Handle(bytes []byte) bool {
	// TODO: send feedback to rabbitmq via some key

	reqObj := dto.SendNotificationRequest{}
	if !layer.JSONBytesConverterMiddleware(bytes, &reqObj) {
		return true
	}

	templateEntity, status := bcnc.templateRepository.Get(reqObj.TemplateID)
	if status == code.StatusError {
		return false
	} else if status == code.StatusNotFound {
		return true
	}

	filledResult := bcnc.fillPlaceholders(templateEntity, &reqObj.UniversalPlaceholders)
	if !filledResult {
		return false
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
			notification := bcnc.toNotificationEntity(currentTarget, templateEntity, templateEntity.Body.Email, &reqObj, &errs)
			if notification != nil {
				go bcnc.sendNotification(*notification, target.Email, bcnc.notificationRepository.SendEmail, &successCount, &failedCount, &wg)
			}
		}


		if target.PhoneNumber != nil {
			notification := bcnc.toNotificationEntity(currentTarget, templateEntity, templateEntity.Body.SMS, &reqObj, &errs)
			if notification != nil {
				go bcnc.sendNotification(*notification, target.PhoneNumber, bcnc.notificationRepository.SendSMS, &successCount, &failedCount, &wg)
			}
		}


		if target.FCMRegistrationToken != nil {
			notification := bcnc.toNotificationEntity(currentTarget, templateEntity, templateEntity.Body.Push, &reqObj, &errs)
			if notification != nil {
				go bcnc.sendNotification(*notification, target.FCMRegistrationToken, bcnc.notificationRepository.SendPush, &successCount, &failedCount, &wg)
			}
		}
	}

	wg.Wait()

	if failedCount > 0 || len(errs) > 0 {
		logrus.Errorf("Failed To Send Notifications: %d\tSucceeded Notifications: %d\tErrors:\n%s", failedCount, successCount, errs)
	}

	return true
}



func (bcnc *basicCreateNotificationV1Controller) sendNotification(
											notification entity.NotificationEntity,
											contactInfo *string,
											outsourceNotification func(notificationEntity *entity.NotificationEntity)bool,
											successCount *int,
											failedCount *int,
											wg *sync.WaitGroup) {
	notification.ContactInfo = *contactInfo

	wg.Add(1)
	defer wg.Done()
	if outsourceNotification(&notification) {
		status := bcnc.notificationRepository.Insert(&notification)
		if status == code.StatusSuccess {
			(*successCount)++
			return
		}
	}
	(*failedCount)++
}

func (bcnc *basicCreateNotificationV1Controller) toNotificationEntity(
											targetId int,
											template *entity.TemplateEntity,
											message *string,
											request *dto.SendNotificationRequest,
											errs *[]dto.SendNotificationErrorData) *entity.NotificationEntity {
	replaced, err := dto.FillPlaceholders(*message, &request.Targets[targetId].Placeholders)
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

func (bcnc *basicCreateNotificationV1Controller) fillPlaceholders(
											template *entity.TemplateEntity,
											placeholders *[]dto.TemplatePlaceholder) bool {
	if template.Body.Email != nil {
		edited, err := dto.FillPlaceholders(*template.Body.Email, placeholders)
		if util.ManageError(err) {
			return false
		}
		template.Body.Email = edited
	}

	if template.Body.SMS != nil {
		edited, err := dto.FillPlaceholders(*template.Body.SMS, placeholders)
		if util.ManageError(err) {
			return false
		}
		template.Body.SMS = edited
	}

	if template.Body.Push != nil {
		edited, err := dto.FillPlaceholders(*template.Body.Push, placeholders)
		if util.ManageError(err) {
			return false
		}
		template.Body.Push = edited
	}

	return true
}