package iface

type IQueryBuilder interface {
	Where(condition string, placeholderValue any) IQueryBuilder
	End(limit *int, offset *int) *string
	Values() *[]any
}