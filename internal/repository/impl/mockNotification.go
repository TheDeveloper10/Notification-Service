package impl

import (
	"notification-service/internal/entity"
	"notification-service/internal/util"
)

type MockNotificationRepository struct {}

func (m MockNotificationRepository) Insert(entity *entity.NotificationEntity) util.RepoStatusCode {
	return util.RepoStatusSuccess
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

func (m MockNotificationRepository) GetBulk(filter *entity.NotificationFilter) (*[]entity.NotificationEntity, util.RepoStatusCode) {
	return &[]entity.NotificationEntity{
		{
			Id: 1,
			TemplateID: 2,
			AppID: "testutils-app",
			ContactInfo: "testutils@example.com",
			Title: "Hi!",
			Message: "Hello!",
			SentTime: 123342352,
		},
	},  util.RepoStatusSuccess
}
