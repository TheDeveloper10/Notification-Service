package util

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type QueryBuilder struct {
	query 	   string
	whereIsSet bool
	placeholderValues *[]any
}

func (qb *QueryBuilder) Core(queryCore string) *QueryBuilder {
	loweredQuery := strings.ToLower(queryCore)

	if strings.Contains(loweredQuery, " where ") || strings.Contains(loweredQuery, " limit ") || strings.Contains(loweredQuery, " offset ") {
		log.Error("Core of query cannot contain 'where', 'limit' or 'offset'!")
		return nil
	}

	qb.query      		 = queryCore
	qb.whereIsSet 		 = false
	qb.placeholderValues = nil

	return qb
}

func (qb *QueryBuilder) Where(condition string, placeholderValue any) *QueryBuilder {
	if qb.placeholderValues == nil {
		qb.placeholderValues = &[]any{ placeholderValue }
	} else {
		v := append(*qb.placeholderValues, placeholderValue)
		qb.placeholderValues = &v
	}
	if qb.whereIsSet {
		qb.query += " and " + condition
	} else {
		qb.query += " where " + condition
		qb.whereIsSet = true
	}

	return qb
}

func (qb *QueryBuilder) End(limit *int, offset *int) *string {
	if limit != nil {
		qb.query += " limit " + strconv.Itoa(*limit)
	}
	if offset != nil {
		qb.query += " offset " + strconv.Itoa(*offset)
	}

	return &qb.query
}

func (qb *QueryBuilder) Values() *[]any {
	return qb.placeholderValues
}