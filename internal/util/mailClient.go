package util

import (
	"fmt"
	"net/smtp"
	"notification-service/internal/util/iface"
	"strconv"
)

type MailClient struct {
	iface.IMailClient
	address   string
	fromEmail string
	auth      smtp.Auth
}

func (mw *MailClient) Init(host string, port int, fromEmail string, password string) {
	mw.address = host + ":" + strconv.Itoa(port)
	mw.fromEmail = fromEmail
	mw.auth = smtp.PlainAuth("", fromEmail, password, host)
}

func (mw *MailClient) Mail(subject string, message string, to []string) error {
	msg := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, message)
	return smtp.SendMail(mw.address, mw.auth, mw.fromEmail, to, []byte(msg))
}

func (mw *MailClient) MailSingle(subject string, message string, to string) error {
	return mw.Mail(subject, message, []string{to})
}
