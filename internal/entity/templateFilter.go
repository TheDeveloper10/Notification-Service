package entity

import (
	"github.com/TheDeveloper10/rem"
	"net/http"
	"notification-service/internal/util"
)

const (
	DefaultTemplatePage = 1
	DefaultTemplateSize = 20
)

type TemplateFilter struct {
	Page  int
	Size  int
}

func TemplateFilterFromRequest(req rem.IRequest, res rem.IResponse) *TemplateFilter {
	filter := TemplateFilter{}
	extractor := util.NewQueryParameterExtractor(req.GetQueryParameters())

	if page, err := extractor.GetPositiveInteger("page", DefaultTemplatePage); err != nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError(err.Error()))
	} else {
		filter.Page = *page
	}

	if size, err := extractor.GetPositiveInteger("size", DefaultTemplateSize); err != nil {
		res.Status(http.StatusBadRequest).JSON(util.ErrorListFromTextError(err.Error()))
	} else {
		filter.Size = *size
	}

	return &filter
}