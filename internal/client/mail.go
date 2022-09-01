package client

import (
	"fmt"
	"net/smtp"
	"notification-service/internal/helper"
	"notification-service/internal/util/iface"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var MailClient iface.IMailClient = nil

func InitializeMailClient() {
	if MailClient != nil {
		return
	}

	if helper.Config.Service.Clients.Has("smtp") {
		client := &mailClient{}
		client.init(helper.Config.SMTP.Host, helper.Config.SMTP.Port, helper.Config.SMTP.FromEmail, helper.Config.SMTP.FromPassword)

		MailClient = client
	} else {
		MailClient = &emptyMailClient{}
	}
}

type mailClient struct {
	iface.IMailClient
	address   string
	fromEmail string
	auth      smtp.Auth
}

func (mw *mailClient) init(host string, port int, fromEmail string, password string) {
	if mw.auth != nil {
		log.Fatal("Cannot initialize a MailClient more than once")
		return
	}
	mw.address = host + ":" + strconv.Itoa(port)
	mw.fromEmail = fromEmail
	mw.auth = smtp.PlainAuth("", fromEmail, password, host)
}

func (mw *mailClient) Mail(subject string, message string, to []string) error {
	msg := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, message)
	return smtp.SendMail(mw.address, mw.auth, mw.fromEmail, to, []byte(msg))
}

func (mw *mailClient) MailSingle(subject string, message string, to string) error {
	return mw.Mail(subject, message, []string{to})
}

type emptyMailClient struct {
	iface.IMailClient
}

func (emc *emptyMailClient) Mail(subject string, message string, to []string) error { return nil }

func (emc *emptyMailClient) MailSingle(subject string, message string, to string) error { return nil }
