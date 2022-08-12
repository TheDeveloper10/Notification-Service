package entity

import (
	"net/http"
	"notification-service/internal/util"
	"notification-service/internal/util/iface"
)

const (
	DefaultTemplatePage = 1
	DefaultTemplateSize = 20
)

type TemplateFilter struct {
	Page  int
	Size  int
}

func TemplateFilterFromRequest(req *http.Request, res iface.IResponseWriter) *TemplateFilter {
	filter := TemplateFilter{}
	extractor := util.NewQueryParameterExtractor(req.URL.Query())

	if page, err := extractor.GetPositiveInteger("page", DefaultTemplatePage); err != nil {
		res.Status(http.StatusBadRequest).Error(err)
	} else {
		filter.Page = *page
	}

	if size, err := extractor.GetPositiveInteger("size", DefaultTemplateSize); err != nil {
		res.Status(http.StatusBadRequest).Error(err)
	} else {
		filter.Size = *size
	}

	return &filter
}