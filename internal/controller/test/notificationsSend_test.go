package test

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/controller"
	"notification-service/internal/controller/layer"
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util/testutils"
	"testing"
)

func TestBasicNotificationV1Controller_Send(t *testing.T) {
	templateRepository := repository.NewMockTemplateRepository()
	notificationRepository := repository.NewMockNotificationRepository()
	clientRepository := repository.NewMockClientRepository()
	tac := controller.NewNotificationV1Controller(templateRepository, notificationRepository, clientRepository)
	router := rem.NewRouter()
	tac.CreateRoutes(router)

	newTestCase := func(reqBody *string, reqHeaders map[string]string, expectedStatusCode int) testutils.ControllerTestCase {
		return testutils.ControllerTestCase{
			Router:          router,
			ReqMethod:       http.MethodPost,
			ReqURL:          "/v1/notifications",
			ReqHeaders:      reqHeaders,
			ReqBody:         reqBody,
			ExpectedStatus:  expectedStatusCode,
		}
	}

	s := func(str string) *string { return &str }

	testCases := []testutils.ControllerTestCase{
		newTestCase(
			layer.ToJSONString(
				&dto.SendNotificationRequest{
					AppID: "testutils",
					TemplateID: 4,
					Title: "Welcome",
					Targets: []dto.NotificationTarget{
						{ Email: s("testutils@example.com"), Placeholders: []dto.TemplatePlaceholder{ { Key: "firstName", Value: "John" } } },
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
				&dto.SendNotificationRequest{
					AppID: "testutils",
					TemplateID: 4,
					Title: "Welcome",
					Targets: []dto.NotificationTarget{
						{ PhoneNumber: s("+357123451234"), Placeholders: []dto.TemplatePlaceholder{ { Key: "firstName", Value: "John" } } },
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
				&dto.SendNotificationRequest{
					AppID: "testutils",
					TemplateID: 3,
					Title: "Welcome",
					Targets: []dto.NotificationTarget{
						{ PhoneNumber: s("+357123451234"), Placeholders: []dto.TemplatePlaceholder{ { Key: "firstName", Value: "John" } } },
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
				&dto.SendNotificationRequest{
					AppID: "testutils",
					TemplateID: 4,
					Title: "Welcome",
					Targets: []dto.NotificationTarget{
						{ FCMRegistrationToken: s("123uji214oiphOUHwouethwoiueth"), Placeholders: []dto.TemplatePlaceholder{ { Key: "firstName", Value: "John" } } },
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
				&dto.SendNotificationRequest{
					AppID: "testutils",
					TemplateID: 2,
					Title: "Welcome",
					Targets: []dto.NotificationTarget{
						{ FCMRegistrationToken: s("123uji214oiphOUHwouethwoiueth"), Placeholders: []dto.TemplatePlaceholder{ { Key: "firstName", Value: "John" } } },
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
				&dto.SendNotificationRequest{
					AppID: "testutils",
					TemplateID: 1,
					Title: "Welcome",
					Targets: []dto.NotificationTarget{
						{ FCMRegistrationToken: s("123uji214oiphOUHwouethwoiueth"), Placeholders: []dto.TemplatePlaceholder{ { Key: "firstName", Value: "John" } } },
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

	testutils.RunControllerTestCases(&testCases, t)
}