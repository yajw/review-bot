package mysql

import (
	"context"
	"errors"
	"sync"
	"time"

	log "github.com/yajw/review-bot/libs/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once

	db *gorm.DB
)

func ConnectDB() {
	once.Do(func() {
		connectDB()
	})
}

func connectDB() {
	dsn := "review_bot:7z$8K@k7@tcp(127.0.0.1:3306)/review?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: &sqlLogger{}})
	if err != nil {
		panic(err)
	}
}

func GetConnection() (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("no connections available")
	}

	return db.Debug(), nil
}

type sqlLogger struct{}

func (s *sqlLogger) LogMode(logger.LogLevel) logger.Interface {
	return s
}

func (s *sqlLogger) Info(ctx context.Context, f string, args ...interface{}) {
	log.Info(f, args)
}

func (s *sqlLogger) Warn(ctx context.Context, f string, args ...interface{}) {
	log.Info(f, args)

}

func (s *sqlLogger) Error(ctx context.Context, f string, args ...interface{}) {
	log.Error(f, args)
}

func (s *sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
}
