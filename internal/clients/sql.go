package clients

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"notification-service.com/packages/internal/helper"
)

func InitializeSQLClient() {
	if SQLClient != nil {
		return
	}

	dbConfig := &helper.Config.Database
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Name)

	db, err := sql.Open(dbConfig.Driver, conn)
	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(dbConfig.PoolSize)
	db.SetMaxOpenConns(dbConfig.PoolSize)
	
	SQLClient = db
}

var SQLClient *sql.DB = nil
