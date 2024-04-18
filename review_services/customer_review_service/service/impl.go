package service

import (
	"github.com/yajw/review-bot/review_services/common"
	"github.com/yajw/review-bot/review_services/ops_review_service/service"
)

type ReviewServiceImpl struct{}

func (r ReviewServiceImpl) GetReviewTemplate(sceneKey string) (*common.ReviewTemplate, error) {
	// a mock implementation here
	return service.GetReviewTemplateBySceneKey(sceneKey), nil
}

func (r ReviewServiceImpl) SubmitReview(req *SubmitReviewRequest) (*SubmitReviewResponse, error) {
	// insert db
	panic("implement me")
}

func (r ReviewServiceImpl) GetUserReviews(userID int64, sceneKey string, sceneIDList []int64) ([]*common.Review, error) {
	// query from cache, if miss query from db
	panic("implement me")
}
