package util

import (
	"notification-service/internal/util/iface"

	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func NewQueryBuilder(queryCore string) iface.IQueryBuilder {
	qb := queryBuilder{}
	qb.initialize(&queryCore)
	return &qb
}

type queryBuilder struct {
	iface.IQueryBuilder
	query 	   string
	whereIsSet bool
	placeholderValues *[]any
}

func (qb *queryBuilder) initialize(queryCore *string) {
	loweredQuery := strings.ToLower(*queryCore)

	if strings.Contains(loweredQuery, " where ") || strings.Contains(loweredQuery, " limit ") || strings.Contains(loweredQuery, " offset ") {
		log.Error("Core of query cannot contain 'where', 'limit' or 'offset'!")
		return
	}

	qb.query      		 = *queryCore
	qb.whereIsSet 		 = false
	qb.placeholderValues = nil
}

func (qb *queryBuilder) Where(condition string, placeholderValue any, skip bool) iface.IQueryBuilder {
	if skip {
		return qb
	}

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

func (qb *queryBuilder) End(limit *int, offset *int) *string {
	if limit != nil {
		qb.query += " limit " + strconv.Itoa(*limit)
	}
	if offset != nil {
		qb.query += " offset " + strconv.Itoa(*offset)
	}

	return &qb.query
}

func (qb *queryBuilder) Values() *[]any {
	return qb.placeholderValues
}