package repository

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	
	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/dto"
)

type NotificationRepository interface {
	Insert(snr *dto.SendNotificationRequest, message *string) bool
}

type basicNotificationRepository struct { }

func NewNotificationRepository() NotificationRepository {
	return &basicNotificationRepository{}
}

func (bnr *basicNotificationRepository) Insert(snr *dto.SendNotificationRequest, message *string) bool {
	stmt, err1 := clients.SQLClient.Prepare("insert into Notifications(Title, ContactType, ContactInfo, Message, UserId, AppId) values(?, ?, ?, ?, ?, ?)")
	if err1 != nil {
		log.Error(err1.Error())
		return false
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(*snr.Title, *snr.ContactType, *snr.ContactInfo, *message, *snr.UserId, *snr.AppId)
	if err2 != nil {
		log.Error(err2.Error())
		return false
	}

	id, err3 := res.LastInsertId()
	if err3 != nil {
		log.Error(err3.Error())
		return false
	}

	log.Info("Inserted template with id " + strconv.FormatInt(id, 10))
	return true
}