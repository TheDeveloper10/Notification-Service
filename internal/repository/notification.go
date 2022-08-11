package repository

import (
	"context"
	"firebase.google.com/go/messaging"
	"notification-service/internal/helper"
	"strconv"

	log "github.com/sirupsen/logrus"

	"notification-service/internal/clients"
	"notification-service/internal/entity"
)

type NotificationRepository interface {
	Insert(entity *entity.NotificationEntity) bool
}

type basicNotificationRepository struct { }

func NewNotificationRepository() NotificationRepository {
	return &basicNotificationRepository{}
}

func (bnr *basicNotificationRepository) Insert(entity *entity.NotificationEntity) bool {
	stmt, err1 := clients.SQLClient.Prepare("insert into Notifications(TemplateId, FCMRegistrationToken, AppId, ContactType, ContactInfo, Title, Message) values(?, ?, ?, ?, ?, ?, ?)")
	if helper.IsError(err1) {
		return false
	}
	defer helper.HandledClose(stmt)

	res1, err2 := stmt.Exec(entity.TemplateID, entity.FCMRegistrationToken, entity.AppID, entity.ContactType, entity.ContactInfo, entity.Title, entity.Message)
	if helper.IsError(err2) {
		return false
	}

	id, err3 := res1.LastInsertId()
	if helper.IsError(err3) {
		return false
	}

	log.Info("Inserted notification into the database with id " + strconv.FormatInt(id, 10))

	_, err := clients.FCMClient.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: entity.Title,
			Body:  entity.Message,
		},
		Token: *entity.FCMRegistrationToken,
	})
	if helper.IsError(err) {
		return false
	}

	log.Info("Sent notification via FCM")
	return true
}