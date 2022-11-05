package client

import (
	"database/sql"
	"fmt"
	"notification-service/internal/util"
	"time"

	"notification-service/internal/util/iface"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var SQLClient iface.ISQLClient = nil

func InitializeSQLClient() {
	if SQLClient != nil {
		return
	}

	dbConfig := &util.Config.Database
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Name)

	client := sqlClient{}
	client.Init(dbConfig.Driver, conn, dbConfig.PoolSize)
	SQLClient = &client
}

type sqlClient struct {
	iface.ISQLClient
	client *sql.DB
}

func (c *sqlClient) Init(driver string, connection string, poolSize int) {
	client, err := sql.Open(driver, connection)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	client.SetConnMaxIdleTime(5 * time.Second)
	client.SetConnMaxLifetime(0)
	client.SetMaxIdleConns(poolSize)
	client.SetMaxOpenConns(poolSize)

	c.client = client
}

func (c *sqlClient) Exec(query string, args ...any) sql.Result {
	stmt, err := c.client.Prepare(query)
	if util.ManageError(err) {
		return nil
	}
	defer util.HandledClose(stmt)

	res, err := stmt.Exec(args...)
	if util.ManageError(err) {
		return nil
	}

	return res
}

func (c *sqlClient) Query(query string, args ...any) *sql.Rows {

	stmt, err := c.client.Prepare(query)
	if util.ManageError(err) {
		return nil
	}
	defer util.HandledClose(stmt)

	rows, err := stmt.Query(args...)
	if util.ManageError(err) {
		return nil
	}

	return rows
}
