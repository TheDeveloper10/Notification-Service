package repository

import (
	"notification-service/internal/util"
	"strconv"

	log "github.com/sirupsen/logrus"

	"notification-service/internal/client"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
)

type NotificationRepository interface {
	Insert(notification *entity.NotificationEntity) bool

	SendEmail(notification *entity.NotificationEntity) bool
	SendPush(notification *entity.NotificationEntity) bool
	SendSMS(notification *entity.NotificationEntity) bool

	GetBulk(filter *entity.NotificationFilter) *[]entity.NotificationEntity
}

type basicNotificationRepository struct{}

func NewNotificationRepository() NotificationRepository {
	return &basicNotificationRepository{}
}

func (bnr *basicNotificationRepository) Insert(notification *entity.NotificationEntity) bool {
	stmt, err1 := client.SQLClient.Prepare("insert into Notifications(AppId, TemplateId, ContactInfo, Title, Message) values(?, ?, ?, ?, ?)")
	if helper.IsError(err1) {
		return false
	}
	defer helper.HandledClose(stmt)

	res1, err2 := stmt.Exec(notification.AppID, notification.TemplateID, notification.ContactInfo, notification.Title, notification.Message)
	if helper.IsError(err2) {
		return false
	}

	id, err3 := res1.LastInsertId()
	if helper.IsError(err3) {
		return false
	}

	log.Info("Inserted notification into the database with id " + strconv.FormatInt(id, 10))
	return true
}

func (bnr *basicNotificationRepository) SendEmail(notification *entity.NotificationEntity) bool {
	err := client.MailClient.MailSingle(notification.Title, notification.Message, notification.ContactInfo)
	if helper.IsError(err) {
		return false
	}

	log.Info("Sent an email")
	return true
}

func (bnr *basicNotificationRepository) SendPush(notification *entity.NotificationEntity) bool {
	err := client.PushClient.SendMessage(notification.Title, notification.Message, notification.ContactInfo)
	if helper.IsError(err) {
		return false
	}

	log.Info("Sent a push notification")
	return true
}

func (bnr *basicNotificationRepository) SendSMS(notification *entity.NotificationEntity) bool {
	return true
}

func (bnr *basicNotificationRepository) GetBulk(filter *entity.NotificationFilter) *[]entity.NotificationEntity {
	builder := util.NewQueryBuilder("select * from Notifications")

	builder.
		Where("AppId=?", filter.AppId, filter.AppId == nil).
		Where("TemplateId=?", filter.TemplateId, filter.TemplateId == nil).
		Where("SentTime>=?", filter.StartTime, filter.StartTime == nil).
		Where("SentTime<=?", filter.EndTime, filter.EndTime == nil)

	offset := (filter.Page - 1) * filter.Size
	query := builder.End(&filter.Size, &offset)
	stmt, err := client.SQLClient.Prepare(*query)
	if helper.IsError(err) {
		return nil
	}

	values := builder.Values()
	if values == nil {
		values = &[]any{}
	}

	rows, err2 := stmt.Query((*values)...)
	if helper.IsError(err2) {
		return nil
	}
	defer helper.HandledClose(rows)

	var notifications []entity.NotificationEntity
	for rows.Next() {
		record := entity.NotificationEntity{}
		err3 := rows.Scan(&record.Id, &record.AppID, &record.TemplateID, &record.ContactInfo,
			&record.Title, &record.Message, &record.SentTime)
		if helper.IsError(err3) {
			return nil
		}
		notifications = append(notifications, record)
	}

	log.Info("Fetched " + strconv.Itoa(len(notifications)) + " template(s)")
	return &notifications
}
