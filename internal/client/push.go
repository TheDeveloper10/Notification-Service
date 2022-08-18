package client

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"notification-service/internal/helper"
	"notification-service/internal/util/iface"

	log "github.com/sirupsen/logrus"
)

var PushClient iface.IPushClient

func InitializePushClient(credentialsFile string) {
	if PushClient != nil {
		return
	}

	if helper.Config.Service.UsePush == "yes" {
		client := &pushClient{}
		client.init(credentialsFile)
		PushClient = client
	} else {
		PushClient = &emptyPushClient{}
	}
}


type pushClient struct {
	iface.IPushClient
	client *messaging.Client
}

func (pc *pushClient) init(credentialsFile string) {
	if pc.client != nil {
		log.Fatal("Cannot initialize a pushClient more than once")
		return
	}

	ctx := context.Background()

	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	pc.client, err = app.Messaging(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (pc *pushClient) SendMessage(title string, body string, token string) error {
	_, err := pc.client.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: token,
	})
	return err
}



type emptyPushClient struct {
	iface.IPushClient
}

func (epc *emptyPushClient) SendMessage(title string, body string, token string) error { return nil }