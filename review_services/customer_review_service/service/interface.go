package service

import (
	"time"

	"github.com/yajw/review-bot/review_services/common"
)

type RateRequest struct {
	UserID int64
	ItemID int64

	Star    int    // 1-5 stars
	Comment string // user left some comment here
}

type RateResponse struct{}

type GetUserRatesRequest struct {
	UserID int64

	ItemIDList []int64
}

type RateInfo struct {
	ID         int64
	UserID     int64
	ItemID     int64
	Star       int
	Comment    string
	CreateTime time.Time
}

type GetUserRatesResponse struct {
	UserRates map[int64]*RateInfo
}

type SubmitReviewRequest struct {
	SceneKey      string                 `json:"scene_key,omitempty"`
	SceneID       int64                  `json:"scene_id,omitempty"`
	UserID        int64                  `json:"user_id,omitempty"`
	ReviewContent string                 `json:"review_content,omitempty"`
	SubmitTime    time.Time              `json:"submit_time"`
	ExtraAttrs    map[string]interface{} `json:"extra_attrs,omitempty"`
}

type SubmitReviewResponse struct{}

// ReviewService is a RPC service for review
type ReviewService interface {
	// SubmitReview user submit a review
	SubmitReview(req *SubmitReviewRequest) (*SubmitReviewResponse, error)

	// GetUserReviews user's review list
	GetUserReviews(uid int64, sceneKey string, sceneIDList []int64) ([]*common.Review, error)

	GetReviewTemplate(sceneKey string) (*common.ReviewTemplate, error)
}
