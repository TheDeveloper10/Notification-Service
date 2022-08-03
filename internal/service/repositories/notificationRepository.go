package repositories

import (
	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/service/dtos"
)

func InsertNotification(snr *dtos.SendNotificationRequest, message *string) bool {
	client := clients.GetMysqlClient()

	stmt, err1 := client.Prepare("insert into Notifications(Title, ContactType, ContactInfo, Message, UserId, AppId) values(?, ?, ?, ?, ?, ?)")
	if err1 != nil {
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(*snr.Title, snr.ContactTypeId(), *snr.ContactInfo, *message, *snr.UserId, *snr.AppId)
	return err2 == nil
}