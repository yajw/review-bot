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

	// index: (scene_key, scene_id)
	// index: (uid, scene_key)
	_, err = db.Exec("create table if not exists user_rates (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"uid int, " +
		"scene_key varchar(255), " +
		"scene_id int, " +
		"review varchar(256), " +
		"attrs text, " +
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
