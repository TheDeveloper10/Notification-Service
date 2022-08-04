package repository

import (
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
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(*snr.Title, *snr.ContactType, *snr.ContactInfo, *message, *snr.UserId, *snr.AppId)
	return err2 == nil
}