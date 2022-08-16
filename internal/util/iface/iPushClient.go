package iface

type IPushClient interface {
	SendMessage(title string, body string, token string) error
}