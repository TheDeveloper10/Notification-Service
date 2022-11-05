package repository

import (
	entity2 "notification-service/internal/data/entity"
	"notification-service/internal/repository/impl"
	"notification-service/internal/util/code"
)

type INotificationRepository interface {
	Insert(*entity2.NotificationEntity) code.StatusCode

	SendEmail(*entity2.NotificationEntity) bool
	SendPush(*entity2.NotificationEntity) bool
	SendSMS(*entity2.NotificationEntity) bool

	GetBulk(*entity2.NotificationFilter) (*[]entity2.NotificationEntity, code.StatusCode)
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