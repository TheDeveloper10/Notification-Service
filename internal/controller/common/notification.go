package common

import (
	"errors"
	"net/http"
	"notification-service/internal/data/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util/code"
	"sync"
)

type NotificationSender struct {
	TemplateRepository     repository.ITemplateRepository
	NotificationRepository repository.INotificationRepository
}

func (ns *NotificationSender) Send(request *dto.SendNotificationRequest) (any, bool, int) {
	templateEntity, status := ns.TemplateRepository.Get(request.TemplateID)
	if status == code.StatusError {
		return nil, false, http.StatusBadRequest
	} else if status == code.StatusNotFound {
		return errors.New("Template not found"), true, http.StatusNotFound
	}

	err := FillPlaceholdersOnTemplate(templateEntity, &request.UniversalPlaceholders)
	if err != nil {
		return err, true, http.StatusUnprocessableEntity
	}

	var wg sync.WaitGroup
	errs := make([]dto.SendNotificationErrorData, 0)

	successCount := 0
	failedCount := 0
	targetCount := len(request.Targets)
	for currentTarget := 0; currentTarget < targetCount; currentTarget++ {
		target := &(request.Targets[currentTarget])

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
			notification := ToNotificationEntity(currentTarget, templateEntity, templateEntity.Body.Email, request, &errs)
			if notification != nil {
				go SendNotification(*notification, target.Email, ns.NotificationRepository.SendEmail, &successCount, &failedCount, ns.NotificationRepository, &wg)
			}
		}


		if target.PhoneNumber != nil {
			notification := ToNotificationEntity(currentTarget, templateEntity, templateEntity.Body.SMS, request, &errs)
			if notification != nil {
				go SendNotification(*notification, target.PhoneNumber, ns.NotificationRepository.SendSMS, &successCount, &failedCount, ns.NotificationRepository, &wg)
			}
		}


		if target.FCMRegistrationToken != nil {
			notification := ToNotificationEntity(currentTarget, templateEntity, templateEntity.Body.Push, request, &errs)
			if notification != nil {
				go SendNotification(*notification, target.FCMRegistrationToken, ns.NotificationRepository.SendPush, &successCount, &failedCount, ns.NotificationRepository, &wg)
			}
		}
	}

	wg.Wait()

	if failedCount > 0 || len(errs) > 0 {
		return &dto.SendNotificationsError{
			Errors:                        errs,
			SuccessfullySentNotifications: successCount,
			FailedNotifications:           failedCount,
		}, true, http.StatusBadRequest
	}

	return nil, true, http.StatusOK
}