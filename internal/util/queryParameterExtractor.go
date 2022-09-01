package util

import (
	"github.com/TheDeveloper10/rem"
	"notification-service/internal/util/iface"

	"errors"
	"strconv"
)

func NewQueryParameterExtractor(values rem.KeyValues) iface.IQueryParameterExtractor {
	return &queryParameterExtractor{
		values: &values,
	}
}

type queryParameterExtractor struct {
	iface.IQueryParameterExtractor
	values *rem.KeyValues
}

func (qpe *queryParameterExtractor) GetPositiveInteger(key string, defaultValue int) (*int, error) {
	if valueStr := qpe.values.Get(key); valueStr != "" {
		valueInt, err := strconv.Atoi(valueStr)
		if err != nil || valueInt <= 0 {
			return nil, errors.New("'" + key + "' must be a positive integer")
		}
		return &valueInt, nil
	}
	return &defaultValue, nil
}

func (qpe *queryParameterExtractor) GetInteger(key string) (*int, error) {
	if valueStr := qpe.values.Get(key); valueStr != "" {
		valueInt, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil, errors.New("'" + key + "' must be a number")
		}
		return &valueInt, nil
	}
	return nil, nil
}

func (qpe *queryParameterExtractor) GetString(key string) *string {
	if str := qpe.values.Get(key); str != "" {
		return &str
	}
	return nil
}