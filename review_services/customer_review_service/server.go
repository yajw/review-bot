package customer_review_service

import (
	"github.com/yajw/review-bot/infra/mysql"
	"github.com/yajw/review-bot/review_services/customer_review_service/service"
)

func Start() {
	mysql.StartServer()
}

func GetReviewService() service.ReviewService {
	return &service.ReviewServiceImpl{}
}
