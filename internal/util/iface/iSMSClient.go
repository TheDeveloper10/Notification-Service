package iface

type ISMSClient interface {
	SendSMS(title string, body string, to string) error
}
