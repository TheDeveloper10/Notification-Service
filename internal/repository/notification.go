package repository

import (
	"notification-service/internal/entity"
	"notification-service/internal/repository/impl"
)

type INotificationRepository interface {
	Insert(*entity.NotificationEntity) bool

	SendEmail(*entity.NotificationEntity) bool
	SendPush(*entity.NotificationEntity) bool
	SendSMS(*entity.NotificationEntity) bool

	GetBulk(*entity.NotificationFilter) *[]entity.NotificationEntity
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