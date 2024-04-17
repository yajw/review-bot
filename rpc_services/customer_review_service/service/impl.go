package service

import (
	"time"

	"github.com/yajw/review-bot/libs/logger"
	"github.com/yajw/review-bot/rpc_services/customer_review_service/cache"
	"github.com/yajw/review-bot/rpc_services/customer_review_service/db"
)

type ReviewServiceImpl struct{}

func (r ReviewServiceImpl) Rate(req *RateRequest) (*RateResponse, error) {
	err := db.CreateRateRecord(req.UserID, req.ItemID, req.Star, req.Comment)
	if err != nil {
		logger.Info("create rate record err: %v", err)
		return nil, err
	}
	return &RateResponse{}, nil
}

func (r ReviewServiceImpl) GetUserRates(req *GetUserRatesRequest) (*GetUserRatesResponse, error) {
	userRates := make(map[int64]*RateInfo)

	missedItemIDs := make([]int64, 0)
	for _, itemID := range req.ItemIDList {
		record := cache.GetRateRecord(req.UserID, itemID)
		if record == nil {
			missedItemIDs = append(missedItemIDs, itemID)
		} else {
			userRates[itemID] = conv(record)
		}
	}

	records, err := db.GetUserRateByUIDAndItemID(req.UserID, missedItemIDs)
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		cache.SetRateRecord(record)
		userRates[record.ItemID] = conv(record)
	}

	return &GetUserRatesResponse{UserRates: userRates}, nil
}

func conv(record *db.RateRecord) *RateInfo {
	return &RateInfo{
		ID:         record.ID,
		UserID:     record.UserID,
		ItemID:     record.ItemID,
		Star:       record.Star,
		Comment:    record.Comment,
		CreateTime: time.Unix(record.CreateTime, 0),
	}
}
