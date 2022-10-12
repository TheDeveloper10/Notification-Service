package impl

import (
	"github.com/sirupsen/logrus"
	"notification-service/internal/client"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
	"notification-service/internal/util"
	"strconv"
)

type BasicNotificationRepository struct{}

func (bnr *BasicNotificationRepository) Insert(notification *entity.NotificationEntity) util.RepoStatusCode {
	res := client.SQLClient.Exec(
		"insert into Notifications(AppId, TemplateId, ContactInfo, Title, Message) values(?, ?, ?, ?, ?)",
		notification.AppID, notification.TemplateID, notification.ContactInfo, notification.Title, notification.Message,
	)
	if res == nil {
		return util.RepoStatusError
	}

	id, err := res.LastInsertId()
	if helper.IsError(err) {
		return util.RepoStatusError
	}

	logrus.Info("Inserted notification into the database with id " + strconv.FormatInt(id, 10))
	return util.RepoStatusSuccess
}

func (bnr *BasicNotificationRepository) SendEmail(notification *entity.NotificationEntity) bool {
	err := client.MailClient.MailSingle(notification.Title, notification.Message, notification.ContactInfo)
	if helper.IsError(err) {
		return false
	}

	logrus.Info("Sent an email")
	return true
}

func (bnr *BasicNotificationRepository) SendPush(notification *entity.NotificationEntity) bool {
	err := client.PushClient.SendMessage(notification.Title, notification.Message, notification.ContactInfo)
	if helper.IsError(err) {
		return false
	}

	logrus.Info("Sent a push notification")
	return true
}

func (bnr *BasicNotificationRepository) SendSMS(notification *entity.NotificationEntity) bool {
	err := client.SMSClient.SendSMS(notification.Title, notification.Message, notification.ContactInfo)
	if helper.IsError(err) {
		return false
	}

	logrus.Info("Sent an SMS")
	return true
}

func (bnr *BasicNotificationRepository) GetBulk(filter *entity.NotificationFilter) (*[]entity.NotificationEntity, util.RepoStatusCode) {
	builder := util.NewQueryBuilder("select * from Notifications")

	builder.
		Where("AppId=?", filter.AppId, filter.AppId == nil).
		Where("TemplateId=?", filter.TemplateId, filter.TemplateId == nil).
		Where("SentTime>=?", filter.StartTime, filter.StartTime == nil).
		Where("SentTime<=?", filter.EndTime, filter.EndTime == nil)

	offset := (filter.Page - 1) * filter.Size
	query := builder.End(&filter.Size, &offset)

	values := builder.Values()
	if values == nil {
		values = &[]any{}
	}

	rows := client.SQLClient.Query(*query, (*values)...)
	if rows == nil {
		return nil, util.RepoStatusError
	}
	defer helper.HandledClose(rows)

	var notifications []entity.NotificationEntity
	for rows.Next() {
		record := entity.NotificationEntity{}
		err3 := rows.Scan(&record.Id, &record.AppID, &record.TemplateID, &record.ContactInfo,
			&record.Title, &record.Message, &record.SentTime)
		if helper.IsError(err3) {
			return nil, util.RepoStatusError
		}
		notifications = append(notifications, record)
	}

	logrus.Info("Fetched " + strconv.Itoa(len(notifications)) + " template(s)")
	return &notifications, util.RepoStatusSuccess
}
