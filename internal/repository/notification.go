package repository

import (
	"context"
	"strconv"

	"firebase.google.com/go/messaging"
	log "github.com/sirupsen/logrus"

	"notification-service/internal/clients"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
)

type NotificationRepository interface {
	Insert(entity *entity.NotificationEntity) bool
}

type basicNotificationRepository struct { }

func NewNotificationRepository() NotificationRepository {
	return &basicNotificationRepository{}
}

func (bnr *basicNotificationRepository) Insert(notification *entity.NotificationEntity) bool {
	stmt, err1 := clients.SQLClient.Prepare("insert into Notifications(TemplateId, , AppId, ContactType, ContactInfo, Title, Message) values(?, ?, ?, ?, ?, ?, ?)")
	if helper.IsError(err1) {
		return false
	}
	defer helper.HandledClose(stmt)

	res1, err2 := stmt.Exec(notification.TemplateID, notification.AppID, notification.ContactType, notification.ContactInfo, notification.Title, notification.Message)
	if helper.IsError(err2) {
		return false
	}

	id, err3 := res1.LastInsertId()
	if helper.IsError(err3) {
		return false
	}

	log.Info("Inserted notification into the database with id " + strconv.FormatInt(id, 10))

	if notification.ContactType == entity.ContactTypeSMS {
		_, err := clients.FCMClient.Send(context.Background(), &messaging.Message{
			Notification: &messaging.Notification{
				Title: notification.Title,
				Body:  notification.Message,
			},
			Token: notification.ContactInfo,
		})
		if helper.IsError(err) {
			return false
		}

		log.Info("Sent notification via FCM")
	}

	return true
}