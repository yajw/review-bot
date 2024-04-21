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

type UserReview struct {
	ID            int64     `gorm:"id"`
	SceneKey      string    `gorm:"scene_key,omitempty"`
	SceneID       int64     `gorm:"scene_id,omitempty"`
	UID           int64     `gorm:"uid,omitempty"`
	ReviewContent string    `gorm:"review_content,omitempty"`
	SubmitTime    time.Time `gorm:"submit_time"`
	ExtraAttrs    string    `gorm:"extra_attrs,omitempty"`
	CreateTime    time.Time `gorm:"create_time"`
	ModifyTime    time.Time `gorm:"modify_time"`
}

func (*UserReview) TableName() string {
	return "user_review"
}

func CreateReviewRecord(record *UserReview) (int64, error) {
	db, err := mysql.GetConnection()
	if err != nil {
		return 0, err
	}

	res := db.Create(record)
	if res.Error != nil {
		logger.Info("create review record err: %v", err)
		return 0, res.Error
	}

	return record.ID, nil
}

func GetUserReviewByUIDAndScene(uid int64, sceneKey string, sceneIDs []int64) ([]*UserReview, error) {
	if len(sceneIDs) == 0 {
		return nil, nil
	}

	if len(sceneIDs) > QueryLimit {
		return nil, errors.New("query size over limit")
	}

	db, err := mysql.GetConnection()
	if err != nil {
		panic(err)
	}

	var records []*UserReview
	err = db.Table("user_review").
		Where("uid = ? and scene_key = ? and scene_id in (?)", uid, sceneKey, sceneIDs).
		Find(&records).Error

	if err != nil {
		logger.Info("query rows err: %v", err)
		return nil, err
	}

	return records, nil
}
