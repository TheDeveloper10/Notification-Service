package test

import (
	"notification-service/internal/controller/common"
	"notification-service/internal/data/entity"
	"notification-service/internal/repository"
	"sync"
	"testing"
)

func TestSendNotification(t *testing.T) {
	successOutsourceNotification := func(notificationEntity *entity.NotificationEntity)bool { return true }
	failedOutsourceNotification := func(notificationEntity *entity.NotificationEntity)bool { return false }

	tests := []struct {
		expectedSuccess int
		expectedFailed int
		notification entity.NotificationEntity
		contactInfo string
		outsourceNotification func(notificationEntity *entity.NotificationEntity)bool
	} {
		{
			expectedSuccess: 1,
			expectedFailed: 0,
			notification: entity.NotificationEntity{},
			contactInfo: "",
			outsourceNotification: successOutsourceNotification,
		},
		{
			expectedSuccess: 0,
			expectedFailed: 1,
			notification: entity.NotificationEntity{ Id: 1 },
			contactInfo: "",
			outsourceNotification: successOutsourceNotification,
		},
		{
			expectedSuccess: 0,
			expectedFailed: 1,
			notification: entity.NotificationEntity{},
			contactInfo: "",
			outsourceNotification: failedOutsourceNotification,
		},
		{
			expectedSuccess: 0,
			expectedFailed: 1,
			notification: entity.NotificationEntity{ Id: 1 },
			contactInfo: "",
			outsourceNotification: failedOutsourceNotification,
		},
	}

	wg := sync.WaitGroup{}
	notificationRepository := repository.NewMockNotificationRepository()
	for testId, test := range tests {
		sc := 0
		fc := 0
		common.SendNotification(
			&test.notification,
			&test.contactInfo,
			test.outsourceNotification,
			&sc,
			&fc,
			notificationRepository,
			&wg,
		)

		if sc != test.expectedSuccess || fc != test.expectedFailed {
			t.Errorf(
				"Test: %d\tExpected Success/Failed: %d/%d\tReceived Success/Failed: %d/%d",
				testId, test.expectedSuccess, test.expectedFailed, sc, fc,
			)
		}
	}
}