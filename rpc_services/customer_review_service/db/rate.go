package db

import (
	"errors"
	"time"

	"github.com/yajw/review-bot/infra/mysql"
	"github.com/yajw/review-bot/libs/logger"
)

const (
	QueryLimit = 1000
)

type RateRecord struct {
	ID         int64  `json:"id,omitempty"`
	UserID     int64  `json:"user_id,omitempty"`
	ItemID     int64  `json:"item_id,omitempty"`
	Star       int    `json:"star,omitempty"`
	Comment    string `json:"comment,omitempty"`
	CreateTime int64  `json:"create_time"`
}

func CreateRateRecord(uid int64, itemID int64, star int, comment string) error {
	db, err := mysql.GetConnection()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		"insert into user_rates (uid, item_id, start, comment, create_time) values (?, ?, ?, ?, ?)",
		uid, itemID, star, comment, time.Now().Unix(),
	)

	if err != nil {
		logger.Info("create rate record err: %v", err)
		return err
	}

	return nil
}

func GetUserRateByUIDAndItemID(uid int64, itemIDs []int64) ([]*RateRecord, error) {
	if len(itemIDs) == 0 {
		return nil, nil
	}

	if len(itemIDs) > QueryLimit {
		return nil, errors.New("query size over limit")
	}

	db, err := mysql.GetConnection()
	if err != nil {
		panic(err)
	}

	var records []*RateRecord
	rows, err := db.Query("select * from user_rates where uid = ? and item_id in (?)", uid, itemIDs)
	if err != nil {
		logger.Info("query rows err: %v", err)
		return nil, err
	}

	_ = rows.Scan(&records)

	return records, nil
}
