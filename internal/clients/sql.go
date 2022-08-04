package clients

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeSQLClient() {
	if SQLClient != nil {
		return
	}

	db, err := sql.Open(os.Getenv("db.driver"), os.Getenv("db.conn"))
	if err != nil {
		panic(err.Error())
	}

	poolSize, _ := strconv.Atoi(os.Getenv("db.pool_size"))
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(poolSize)
	db.SetMaxOpenConns(poolSize)
	
	SQLClient = db
}

var SQLClient *sql.DB = nil
