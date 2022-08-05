package repository

import (
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
	if err1 != nil {
		log.Error(err1.Error())
		return false
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(entity.TemplateID, entity.UserID, entity.AppID, entity.ContactType, entity.ContactInfo, entity.Title, entity.Message)
	if err2 != nil {
		log.Error(err2.Error())
		return false
	}

	id, err3 := res.LastInsertId()
	if err3 != nil {
		log.Error(err3.Error())
		return false
	}

	log.Info("Inserted notification with id " + strconv.FormatInt(id, 10))
	return true
}