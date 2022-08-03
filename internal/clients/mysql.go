package clients

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func createClient() (*sql.DB) {
	db, err := sql.Open(os.Getenv("db.driver"), os.Getenv("db.conn"))
	if err != nil {
		panic(err.Error())
	}

	poolSize, _ := strconv.Atoi(os.Getenv("db.pool_size"))
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(poolSize)
	db.SetMaxOpenConns(poolSize)
	
	return db
}

var client *sql.DB = nil

func GetMysqlClient() (*sql.DB) {
	if client == nil {
		client = createClient()
	}
	
	return client
}