package impl

import (
	"github.com/sirupsen/logrus"
	"notification-service/internal/client"
	entity2 "notification-service/internal/data/entity"
	"notification-service/internal/util"
	"notification-service/internal/util/code"
	"strconv"
)

type BasicNotificationRepository struct{}

func (bnr *BasicNotificationRepository) Insert(notification *entity2.NotificationEntity) code.StatusCode {
	res := client.SQLClient.Exec(
		"insert into Notifications(AppId, TemplateId, ContactInfo, Title, Message) values(?, ?, ?, ?, ?)",
		notification.AppID, notification.TemplateID, notification.ContactInfo, notification.Title, notification.Message,
	)
	if res == nil {
		return code.StatusError
	}

	id, err := res.LastInsertId()
	if util.ManageError(err) {
		return code.StatusError
	}

	logrus.Info("Inserted notification into the database with id " + strconv.FormatInt(id, 10))
	return code.StatusSuccess
}

func (bnr *BasicNotificationRepository) SendEmail(notification *entity2.NotificationEntity) bool {
	err := client.MailClient.MailSingle(notification.Title, notification.Message, notification.ContactInfo)
	if util.ManageError(err) {
		return false
	}

	logrus.Info("Sent an email")
	return true
}

func (bnr *BasicNotificationRepository) SendPush(notification *entity2.NotificationEntity) bool {
	err := client.PushClient.SendMessage(notification.Title, notification.Message, notification.ContactInfo)
	if util.ManageError(err) {
		return false
	}

	logrus.Info("Sent a push notification")
	return true
}

func (bnr *BasicNotificationRepository) SendSMS(notification *entity2.NotificationEntity) bool {
	err := client.SMSClient.SendSMS(notification.Title, notification.Message, notification.ContactInfo)
	if util.ManageError(err) {
		return false
	}

	logrus.Info("Sent an SMS")
	return true
}

func (bnr *BasicNotificationRepository) GetBulk(filter *entity2.NotificationFilter) (*[]entity2.NotificationEntity, code.StatusCode) {
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
		return nil, code.StatusError
	}
	defer util.HandledClose(rows)

	var notifications []entity2.NotificationEntity
	for rows.Next() {
		record := entity2.NotificationEntity{}
		err3 := rows.Scan(&record.Id, &record.AppID, &record.TemplateID, &record.ContactInfo,
			&record.Title, &record.Message, &record.SentTime)
		if util.ManageError(err3) {
			return nil, code.StatusError
		}
		notifications = append(notifications, record)
	}

	logrus.Info("Fetched " + strconv.Itoa(len(notifications)) + " template(s)")
	return &notifications, code.StatusSuccess
}
