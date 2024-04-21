package service

import (
	"time"

	"github.com/yajw/review-bot/review_services/common"
)

type SubmitReviewRequest struct {
	SceneKey      string                 `json:"scene_key,omitempty"`
	SceneID       int64                  `json:"scene_id,omitempty"`
	UserID        int64                  `json:"user_id,omitempty"`
	ReviewContent string                 `json:"review_content,omitempty"`
	SubmitTime    time.Time              `json:"submit_time"`
	ExtraAttrs    map[string]interface{} `json:"extra_attrs,omitempty"`
}

type SubmitReviewResponse struct {
	ReviewID int64 `json:"review_id"`
}

// ReviewService is a RPC service for review
type ReviewService interface {
	// SubmitReview user submit a review
	SubmitReview(req *SubmitReviewRequest) (*SubmitReviewResponse, error)

	// GetUserReviews user's review list
	GetUserReviews(uid int64, sceneKey string, sceneIDList []int64) ([]*common.Review, error)

	GetReviewTemplate(sceneKey string) (*common.ReviewTemplate, error)
}
