package clients

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"notification-service/internal/helper"
)

var SQLClient *sql.DB = nil

func InitializeSQLClient() {
	if SQLClient != nil {
		return
	}

	dbConfig := &helper.Config.Database
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Name)

	client, err := sql.Open(dbConfig.Driver, conn)
	if err != nil {
		log.Fatal(err.Error())
	}

	client.SetConnMaxIdleTime(5 * time.Second)
	client.SetConnMaxLifetime(0)
	client.SetMaxIdleConns(dbConfig.PoolSize)
	client.SetMaxOpenConns(dbConfig.PoolSize)
	
	SQLClient = client
}
