package test

import (
	"encoding/json"
	"notification-service/internal/util/iface"
	"reflect"
	"testing"
)

type ControllerTestCase struct {
	Controller iface.IRabbitMQController

	Body any
	ExpectedAck bool

	NoResponseComparison bool
	ExpectedResponse any
}

func (ctc *ControllerTestCase) RunTest(testId int, t *testing.T) {
	data, _ := json.Marshal(ctc.Body)
	response, acknowledge := ctc.Controller.Handle(data)

	if acknowledge != ctc.ExpectedAck {
		t.Errorf("Test %d\tExpected Ack: %t\tReceived: %t", testId, ctc.ExpectedAck, acknowledge)
	}

	if ctc.NoResponseComparison {
		return
	}
	if !reflect.DeepEqual(response, ctc.ExpectedResponse) {
		t.Errorf("Test %d\tExpected Body: %+v\tReceived Body: %+v", testId, ctc.ExpectedResponse, response)
	}
}

func RunControllerTestCases(cases *[]ControllerTestCase, t *testing.T) {
	for testId, testCase := range *cases {
		testCase.RunTest(testId, t)
	}
}
