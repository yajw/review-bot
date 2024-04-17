package customer_review_service

import (
	"github.com/yajw/review-bot/infra/mysql"
	"github.com/yajw/review-bot/rpc_services/customer_review_service/cache"
	"github.com/yajw/review-bot/rpc_services/customer_review_service/service"
)

func Start() {
	cache.InitCache()
	mysql.StartServer()
}

func GetReviewService() service.ReviewService {
	return &service.ReviewServiceImpl{}
}
