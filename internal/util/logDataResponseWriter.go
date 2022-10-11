package util

import (
	"encoding/json"
	"github.com/TheDeveloper10/rem"
	"github.com/sirupsen/logrus"
)

type LogDataResponseWriter struct {
	rem.IResponse
	StatusCode *int
}

func (sorw *LogDataResponseWriter) Status(statusCode int) rem.IResponse {
	sorw.StatusCode = &statusCode
	return sorw
}

func (sorw *LogDataResponseWriter) Header(key string, value string) rem.IResponse { return sorw }
func (sorw *LogDataResponseWriter) Bytes(data []byte) rem.IResponse               { return sorw }
func (sorw *LogDataResponseWriter) Text(text string) rem.IResponse {
	logrus.Info("(LDRW) " + text)
	return sorw
}
func (sorw *LogDataResponseWriter) JSON(data interface{}) rem.IResponse {
	bytes, err := json.Marshal(data)
	if err != nil {
		logrus.Error("(LDRW) Failed to marshal JSON")
		return sorw
	}
	logrus.Info("(LDRW) " + string(bytes))
	return sorw
}