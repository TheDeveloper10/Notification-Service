package test

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller"
	"notification-service/internal/controller/layer"
	dto2 "notification-service/internal/data/dto"
	"notification-service/internal/repository"
	"testing"
)

func TestBasicNotificationV1Controller_Send(t *testing.T) {
	templateRepository := repository.NewMockTemplateRepository()
	notificationRepository := repository.NewMockNotificationRepository()
	clientRepository := repository.NewMockClientRepository()
	tac := controller.NewNotificationV1Controller(templateRepository, notificationRepository, clientRepository)
	router := rem.NewRouter()
	tac.CreateRoutes(router)

	newTestCase := func(reqBody *string, reqHeaders map[string]string, expectedStatusCode int) ControllerTestCase {
		return ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodPost,
			ReqURL:          "/v1/notifications",
			ReqHeaders:      reqHeaders,
			ReqBody:         reqBody,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	s := func(str string) *string { return &str }

	testCases := []ControllerTestCase{
		newTestCase(
			layer.ToJSONString(
				&dto2.SendNotificationRequest{
					AppID: "test",
					TemplateID: 4,
					Title: "Welcome",
					Targets: []dto2.NotificationTarget{
						{ Email: s("test@example.com"), Placeholders: []dto2.TemplatePlaceholder{{ Key: "firstName", Value: "John" } } },
					},
				},
			),
			map[string]string{
				"Authorization": "Bearer 13124",
				"Content-Type": "application/json",
			},
			http.StatusCreated,
		),

		newTestCase(
			layer.ToJSONString(
				&dto2.SendNotificationRequest{
					AppID: "test",
					TemplateID: 4,
					Title: "Welcome",
					Targets: []dto2.NotificationTarget{
						{ PhoneNumber: s("+357123451234"), Placeholders: []dto2.TemplatePlaceholder{{ Key: "firstName", Value: "John" } } },
					},
				},
			),
			map[string]string{
				"Authorization": "Bearer 13124",
				"Content-Type": "application/json",
			},
			http.StatusBadRequest,
		),

		newTestCase(
			layer.ToJSONString(
				&dto2.SendNotificationRequest{
					AppID: "test",
					TemplateID: 3,
					Title: "Welcome",
					Targets: []dto2.NotificationTarget{
						{ PhoneNumber: s("+357123451234"), Placeholders: []dto2.TemplatePlaceholder{{ Key: "firstName", Value: "John" } } },
					},
				},
			),
			map[string]string{
				"Authorization": "Bearer 13124",
				"Content-Type": "application/json",
			},
			http.StatusCreated,
		),

		newTestCase(
			layer.ToJSONString(
				&dto2.SendNotificationRequest{
					AppID: "test",
					TemplateID: 4,
					Title: "Welcome",
					Targets: []dto2.NotificationTarget{
						{ FCMRegistrationToken: s("123uji214oiphOUHwouethwoiueth"), Placeholders: []dto2.TemplatePlaceholder{{ Key: "firstName", Value: "John" } } },
					},
				},
			),
			map[string]string{
				"Authorization": "Bearer 13124",
				"Content-Type": "application/json",
			},
			http.StatusBadRequest,
		),

		newTestCase(
			layer.ToJSONString(
				&dto2.SendNotificationRequest{
					AppID: "test",
					TemplateID: 2,
					Title: "Welcome",
					Targets: []dto2.NotificationTarget{
						{ FCMRegistrationToken: s("123uji214oiphOUHwouethwoiueth"), Placeholders: []dto2.TemplatePlaceholder{{ Key: "firstName", Value: "John" } } },
					},
				},
			),
			map[string]string{
				"Authorization": "Bearer 13124",
				"Content-Type": "application/json",
			},
			http.StatusCreated,
		),

		newTestCase(
			layer.ToJSONString(
				&dto2.SendNotificationRequest{
					AppID: "test",
					TemplateID: 1,
					Title: "Welcome",
					Targets: []dto2.NotificationTarget{
						{ FCMRegistrationToken: s("123uji214oiphOUHwouethwoiueth"), Placeholders: []dto2.TemplatePlaceholder{{ Key: "firstName", Value: "John" } } },
					},
				},
			),
			map[string]string{
				"Authorization": "Bearer 13124",
				"Content-Type": "application/json",
			},
			http.StatusNotFound,
		),
	}

	RunControllerTestCases(&testCases, t)
}