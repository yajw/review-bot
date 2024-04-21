package service

import (
	"encoding/json"
	"time"

	"github.com/yajw/review-bot/libs/logger"
	"github.com/yajw/review-bot/review_services/common"
	"github.com/yajw/review-bot/review_services/customer_review_service/db"
	"github.com/yajw/review-bot/review_services/ops_review_service/service"
)

type ReviewServiceImpl struct{}

func (r ReviewServiceImpl) GetReviewTemplate(sceneKey string) (*common.ReviewTemplate, error) {
	// a mock implementation here
	return service.GetReviewTemplateBySceneKey(sceneKey), nil
}

func (r ReviewServiceImpl) SubmitReview(req *SubmitReviewRequest) (*SubmitReviewResponse, error) {
	bts, err := json.Marshal(req.ExtraAttrs)
	if err != nil {
		logger.Error("marshal extra attrs err: %+v", err)
		return nil, err
	}

	record := &db.UserReview{
		SceneKey:      req.SceneKey,
		SceneID:       req.SceneID,
		UID:           req.UserID,
		ReviewContent: req.ReviewContent,
		SubmitTime:    req.SubmitTime,
		ExtraAttrs:    string(bts),
		CreateTime:    time.Now(),
		ModifyTime:    time.Now(),
	}

	reviewID, err := db.CreateReviewRecord(record)
	if err != nil {
		return nil, err
	}

	return &SubmitReviewResponse{ReviewID: reviewID}, nil
}

func (r ReviewServiceImpl) GetUserReviews(userID int64, sceneKey string, sceneIDList []int64) ([]*common.Review, error) {
	records, err := db.GetUserReviewByUIDAndScene(userID, sceneKey, sceneIDList)
	if err != nil {
		return nil, err
	}

	reviews := make([]*common.Review, 0, len(records))

	for _, record := range records {
		if record == nil {
			continue
		}

		attrs := make(map[string]interface{})
		err := json.Unmarshal([]byte(record.ExtraAttrs), &attrs)
		if err != nil {
			logger.Error("unmarshal attrs err: %+v", err)
			return nil, err
		}

		review := &common.Review{
			SceneKey:      record.SceneKey,
			SceneID:       record.SceneID,
			UserID:        record.UID,
			ReviewContent: record.ReviewContent,
			SubmitTime:    record.SubmitTime,
			ExtraAttrs:    attrs,
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}
