package impl

import (
	"notification-service/internal/data/entity"
	"notification-service/internal/util/code"
)

type MockNotificationRepository struct {}

func (m MockNotificationRepository) Insert(notification *entity.NotificationEntity) code.StatusCode {
	if notification.Id == 1 {
		return code.StatusError
	}
	return code.StatusSuccess
}

func (m MockNotificationRepository) SendEmail(notification *entity.NotificationEntity) bool {
	return true
}

func (m MockNotificationRepository) SendPush(notification *entity.NotificationEntity) bool {
	return true
}

func (m MockNotificationRepository) SendSMS(notification *entity.NotificationEntity) bool {
	return true
}

func (m MockNotificationRepository) GetBulk(filter *entity.NotificationFilter) (*[]entity.NotificationEntity, code.StatusCode) {
	return &[]entity.NotificationEntity{
		{
			Id: 1,
			TemplateID: 2,
			AppID: "test-app",
			ContactInfo: "test@example.com",
			Title: "Hi!",
			Message: "Hello!",
			SentTime: 123342352,
		},
	}, code.StatusSuccess
}
