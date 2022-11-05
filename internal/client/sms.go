package client

import (
	"errors"
	"fmt"
	"net/http"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
	"strings"

	"github.com/sirupsen/logrus"
)

var SMSClient iface.ISMSClient = nil

func InitializeSMSClient() {
	if SMSClient != nil {
		return
	}

	if util.Config.Service.Clients.Has("sms") {
		client := &smsClient{}
		client.init(util.Config.Twillio.AccountSID, util.Config.Twillio.MessagingServiceSID, util.Config.Twillio.AuthToken)

		SMSClient = client
	} else {
		SMSClient = &emptySMSClient{}
	}
}

type smsClient struct {
	iface.ISMSClient
	httpClient *http.Client
	endpoint   string
	parameters string

	accountSID string
	authToken  string
}

func (sc *smsClient) init(accountSID string, messagingServiceSID string, authToken string) {
	if sc.httpClient != nil {
		logrus.Fatal("Cannot initialize a SMSClient more than once")
		return
	}

	sc.httpClient = &http.Client{}
	sc.endpoint = fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", accountSID)
	sc.parameters = "MessagingServiceSid=" + messagingServiceSID + "&To=%s&Body=%s"
	sc.accountSID = accountSID
	sc.authToken = authToken
}

func (sc *smsClient) SendSMS(title string, body string, to string) error {
	message := fmt.Sprintf("%s\r\n%s", title, body)
	relativeParameters := fmt.Sprintf(sc.parameters, to, message)
	data := strings.NewReader(relativeParameters)

	req, err := http.NewRequest("POST", sc.endpoint, data)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(sc.accountSID, sc.authToken)
	resp, err := sc.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return errors.New("Something went wrong sending the SMS!")
	}

	return nil
}

type emptySMSClient struct {
	iface.IMailClient
}

func (esc *emptySMSClient) SendSMS(title string, body string, to string) error { return nil }
