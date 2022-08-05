package repository

import (
	"notification-service.com/packages/internal/helper"
	"strconv"

	log "github.com/sirupsen/logrus"

	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/entity"
)

type NotificationRepository interface {
	Insert(entity *entity.NotificationEntity) bool
}

type basicNotificationRepository struct { }

func NewNotificationRepository() NotificationRepository {
	return &basicNotificationRepository{}
}

func (bnr *basicNotificationRepository) Insert(entity *entity.NotificationEntity) bool {
	stmt, err1 := clients.SQLClient.Prepare("insert into Notifications(TemplateId, UserId, AppId, ContactType, ContactInfo, Title, Message) values(?, ?, ?, ?, ?, ?, ?)")
	if helper.IsError(err1) {
		return false
	}
	defer helper.HandledClose(stmt)

	res, err2 := stmt.Exec(entity.TemplateID, entity.UserID, entity.AppID, entity.ContactType, entity.ContactInfo, entity.Title, entity.Message)
	if helper.IsError(err2) {
		return false
	}

	id, err3 := res.LastInsertId()
	if helper.IsError(err3) {
		return false
	}

	log.Info("Inserted notification with id " + strconv.FormatInt(id, 10))
	return true
}