package entity

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/util"
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

func NotificationFilterFromRequest(req rem.IRequest, res rem.IResponse) *NotificationFilter {
	filter := NotificationFilter{}
	extractor := util.NewQueryParameterExtractor(req.GetQueryParameters())

	filter.AppId = extractor.GetString("appId")

	if page, err := extractor.GetPositiveInteger("page", DefaultNotificationPage); err != nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError(err.Error()))
	} else {
		filter.Page = *page
	}

	if size, err := extractor.GetPositiveInteger("size", DefaultNotificationSize); err != nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError(err.Error()))
	} else {
		filter.Size = *size
	}

	if templateId, err := extractor.GetInteger("templateId"); err != nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError(err.Error()))
	} else {
		filter.TemplateId = templateId
	}

	if startTime, err := extractor.GetInteger("startTime"); err != nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError(err.Error()))
	} else {
		filter.StartTime = startTime
	}

	if endTime, err := extractor.GetInteger("endTime"); err != nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError(err.Error()))
	} else {
		filter.EndTime = endTime
	}

	return &filter
}
