package impl

import (
	entity2 "notification-service/internal/data/entity"
	"notification-service/internal/util/code"
)

type MockNotificationRepository struct {}

func (m MockNotificationRepository) Insert(entity *entity2.NotificationEntity) code.StatusCode {
	return code.StatusSuccess
}

func (m MockNotificationRepository) SendEmail(entity *entity2.NotificationEntity) bool {
	return true
}

func (m MockNotificationRepository) SendPush(entity *entity2.NotificationEntity) bool {
	return true
}

func (m MockNotificationRepository) SendSMS(entity *entity2.NotificationEntity) bool {
	return true
}

func (m MockNotificationRepository) GetBulk(filter *entity2.NotificationFilter) (*[]entity2.NotificationEntity, code.StatusCode) {
	return &[]entity2.NotificationEntity{
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
