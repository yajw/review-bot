package service

import (
	"time"
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

// ReviewService is a RPC service for review
type ReviewService interface {
	// Rate user request to rate an item
	Rate(req *RateRequest) (*RateResponse, error)
	// GetUserRates return a user's rates for one or more items
	GetUserRates(req *GetUserRatesRequest) (*GetUserRatesResponse, error)
}
