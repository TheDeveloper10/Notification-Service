package iface

import "database/sql"

type ISQLClient interface {
	Exec(query string, args ...any) sql.Result
	Query(query string, args ...any) *sql.Rows
}
