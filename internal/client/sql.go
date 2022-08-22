package client

import (
	"database/sql"
	"fmt"
	"time"

	"notification-service/internal/helper"
	"notification-service/internal/util/iface"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var SQLClient iface.ISQLClient

func InitializeSQLClient() {
	if SQLClient != nil {
		return
	}

	dbConfig := &helper.Config.Database
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Name)

	SQLClient := &sqlClient{}
	SQLClient.Init(dbConfig.Driver, conn, dbConfig.PoolSize)
}

type sqlClient struct {
	iface.ISQLClient
	client *sql.DB
}

func (c *sqlClient) Init(driver string, connection string, poolSize int) {
	client, err := sql.Open(driver, connection)
	if err != nil {
		log.Fatal(err.Error())
	}

	client.SetConnMaxIdleTime(5 * time.Second)
	client.SetConnMaxLifetime(0)
	client.SetMaxIdleConns(poolSize)
	client.SetMaxOpenConns(poolSize)

	c.client = client
}

func (c *sqlClient) Exec(query string, args ...any) *sql.Result {
	stmt, err := c.client.Prepare(query)
	if helper.IsError(err) {
		return nil
	}
	defer helper.HandledClose(stmt)

	res, err := stmt.Exec(args...)
	if helper.IsError(err) {
		return nil
	}

	return &res
}

func (c *sqlClient) Query(query string, args ...any) *sql.Rows {

	stmt, err := c.client.Prepare(query)
	if helper.IsError(err) {
		return nil
	}
	defer helper.HandledClose(stmt)

	rows, err := stmt.Query(args...)
	if helper.IsError(err) {
		return nil
	}

	return rows
}
