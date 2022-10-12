package repository

import (
	"notification-service/internal/entity"
	"notification-service/internal/repository/impl"
	"notification-service/internal/util"
)

type INotificationRepository interface {
	Insert(*entity.NotificationEntity) util.RepoStatusCode

	SendEmail(*entity.NotificationEntity) bool
	SendPush(*entity.NotificationEntity) bool
	SendSMS(*entity.NotificationEntity) bool

	GetBulk(*entity.NotificationFilter) (*[]entity.NotificationEntity, util.RepoStatusCode)
}

// ----------------------------------
// Notification Repository Factories
// ----------------------------------

func NewNotificationRepository() INotificationRepository {
	return &impl.BasicNotificationRepository{}
}

func NewMockNotificationRepository() INotificationRepository {
	return &impl.MockNotificationRepository{}
}