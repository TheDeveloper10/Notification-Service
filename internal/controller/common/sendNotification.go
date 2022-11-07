package common

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/repository"
	"notification-service/internal/util/code"
	"sync"
)

func SendNotification(
					notification entity.NotificationEntity,
					contactInfo *string,
					outsourceNotification func(notificationEntity *entity.NotificationEntity)bool,
					successCount *int,
					failedCount *int,
					notificationRepository repository.INotificationRepository,
					wg *sync.WaitGroup) {
	notification.ContactInfo = *contactInfo

	wg.Add(1)
	defer wg.Done()
	if outsourceNotification(&notification) {
		status := notificationRepository.Insert(&notification)
		if status == code.StatusSuccess {
			(*successCount)++
			return
		}
	}
	(*failedCount)++
}
