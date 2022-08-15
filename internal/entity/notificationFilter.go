package entity

import (
	"net/http"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

const (
	DefaultNotificationPage = 1
	DefaultNotificationSize = 20
)

type NotificationFilter struct {
	AppId      *string
	TemplateId *int
	StartTime  *int
	EndTime    *int
	Page       int
	Size       int
}

func NotificationFilterFromRequest(req *http.Request, res iface.IResponseWriter) *NotificationFilter {
	filter := NotificationFilter{}
	extractor := util.NewQueryParameterExtractor(req.URL.Query())

	filter.AppId = extractor.GetString("appId")

	if page, err := extractor.GetPositiveInteger("page", DefaultNotificationPage); err != nil {
		res.Status(http.StatusBadRequest).Error(err)
	} else {
		filter.Page = *page
	}

	if size, err := extractor.GetPositiveInteger("size", DefaultNotificationSize); err != nil {
		res.Status(http.StatusBadRequest).Error(err)
	} else {
		filter.Size = *size
	}

	if templateId, err := extractor.GetInteger("templateId"); err != nil {
		res.Status(http.StatusBadRequest).Error(err)
	} else {
		filter.TemplateId = templateId
	}

	if startTime, err := extractor.GetInteger("startTime"); err != nil {
		res.Status(http.StatusBadRequest).Error(err)
	} else {
		filter.StartTime = startTime
	}

	if endTime, err := extractor.GetInteger("endTime"); err != nil {
		res.Status(http.StatusBadRequest).Error(err)
	} else {
		filter.EndTime = endTime
	}

	return &filter
}
