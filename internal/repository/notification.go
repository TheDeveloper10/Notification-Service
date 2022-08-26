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

func NewNotificationRepository(isMock bool) INotificationRepository {
	if isMock {
		return &impl.MockNotificationRepository{}
	} else {
		return &impl.BasicNotificationRepository{}
	}
}
