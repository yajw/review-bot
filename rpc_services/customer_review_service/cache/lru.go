package cache

import (
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/yajw/review-bot/rpc_services/customer_review_service/db"
)

var (
	once sync.Once

	userRateRecordCache *expirable.LRU[string, *db.RateRecord]
)

func InitCache() {
	once.Do(func() {
		userRateRecordCache = expirable.NewLRU[string, *db.RateRecord](5, nil, time.Millisecond*10)
	})
}

func SetRateRecord(v *db.RateRecord) {
	userRateRecordCache.Add(buildCacheKey(v.UserID, v.ItemID), v)
}

func GetRateRecord(uid int64, itemID int64) *db.RateRecord {
	v, ok := userRateRecordCache.Get(buildCacheKey(uid, itemID))
	if !ok {
		return nil
	}
	return v
}

func buildCacheKey(uid int64, itemID int64) string {
	return fmt.Sprintf("%d:%d", uid, itemID)
}
