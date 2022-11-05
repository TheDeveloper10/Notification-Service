package client

import (
	"context"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"

	"github.com/sirupsen/logrus"
)

var PushClient iface.IPushClient = nil

func InitializePushClient(credentialsFile string) {
	if PushClient != nil {
		return
	}

	if util.Config.Service.Clients.Has("push") {
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
		logrus.Fatal("Cannot initialize a pushClient more than once")
		return
	}

	ctx := context.Background()

	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

	pc.client, err = app.Messaging(ctx)
	if err != nil {
		logrus.Fatal(err)
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
