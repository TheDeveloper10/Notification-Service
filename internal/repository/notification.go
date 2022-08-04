package repository

import (
	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/dto"
)

type notification struct { }

func NewNotificationRepository() *notification {
	return &notification{}
}

func (n *notification) Insert(snr *dto.SendNotificationRequest, message *string) bool {
	stmt, err1 := clients.SQLClient.Prepare("insert into Notifications(Title, ContactType, ContactInfo, Message, UserId, AppId) values(?, ?, ?, ?, ?, ?)")
	if err1 != nil {
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(*snr.Title, snr.ContactTypeId(), *snr.ContactInfo, *message, *snr.UserId, *snr.AppId)
	return err2 == nil
}