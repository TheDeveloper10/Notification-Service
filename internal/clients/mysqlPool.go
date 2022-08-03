package clients

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	clientPool []client
	poolSize int = -1
)

type client struct {
	Used bool
	Db   *sql.DB
}

func (c* client) Close() {
	c.Used = false
}

func createClient() (*sql.DB) {
	db, err := sql.Open("mysql", os.Getenv("db.conn"))
	if err != nil {
		panic(err.Error())
	}
	return db
}

func initPool() {
	if poolSize >= 0 { 
		return 
	}

	poolSizeEnv, _ := strconv.Atoi(os.Getenv("db.pool_size"))
	poolSize = poolSizeEnv

	clientPool = make([]client, poolSize)
	for i := 0; i < poolSize; i++ {
		clientPool[i] = client {
			Used: false,
			Db: createClient(),
		}
	}
}

func GetMysqlClient() (*client) {
	initPool()

	for {
		for i := 0; i < poolSize; i++ {
			if clientPool[i].Used {
				continue
			}
			clientPool[i].Used = true
			return &(clientPool[i])
		}

		time.Sleep(30 * time.Millisecond)
	}
}