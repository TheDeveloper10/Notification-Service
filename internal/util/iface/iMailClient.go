package iface

type IMailClient interface {
	Init(host string, port int, fromEmail string, password string)
	Mail(subject string, message string, to []string) error
	MailSingle(subject string, message string, to string) error
}
