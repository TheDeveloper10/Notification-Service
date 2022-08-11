package clients

import (
	"context"
	log "github.com/sirupsen/logrus"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var FCMClient *messaging.Client = nil

func InitializeFCMClient(credentialsFile string) {
	if FCMClient != nil {
		return
	}

	ctx := context.Background()

	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	FCMClient = client
}
