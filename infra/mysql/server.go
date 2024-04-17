package mysql

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/yajw/review-bot/libs/logger"
)

var (
	once sync.Once

	db *sql.DB
)

func StartServer() {
	once.Do(func() {
		startServer()
	})
}

func startServer() {
	var err error
	db, err = sql.Open("sqlite3", "file:review.db?cache=shared")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(1)

	_, err = db.Exec("create table if not exists user_rates (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"uid int, " +
		"item_id int, " +
		"comment varchar(256), " +
		"star int, " +
		"create_time int, " +
		")",
	)

	if err != nil {
		logger.Info("create table err: %v", err)
		panic(err)
	}
}

func GetConnection() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New("no connections available")
	}

	return db, nil
}
