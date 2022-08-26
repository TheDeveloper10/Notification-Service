package impl

import (
	"notification-service/internal/entity"
)

type MockNotificationRepository struct {}

func (m MockNotificationRepository) Insert(entity *entity.NotificationEntity) bool {
	return true
}

func (m MockNotificationRepository) SendEmail(entity *entity.NotificationEntity) bool {
	return true
}

func (m MockNotificationRepository) SendPush(entity *entity.NotificationEntity) bool {
	return true
}

func (m MockNotificationRepository) SendSMS(entity *entity.NotificationEntity) bool {
	return true
}

func (m MockNotificationRepository) GetBulk(filter *entity.NotificationFilter) *[]entity.NotificationEntity {
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
	}
}
