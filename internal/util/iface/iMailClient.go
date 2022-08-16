package iface

type IMailClient interface {
	Mail(subject string, message string, to []string) error
	MailSingle(subject string, message string, to string) error
}
