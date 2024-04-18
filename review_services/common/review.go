package common

import (
	"time"
)

type Review struct {
	ID            int64                  `json:"id,omitempty"`
	SceneKey      string                 `json:"scene_key,omitempty"`
	SceneID       int64                  `json:"scene_id,omitempty"`
	UserID        int64                  `json:"user_id,omitempty"`
	ReviewContent string                 `json:"review_content,omitempty"`
	SubmitTime    time.Time              `json:"submit_time"`
	ExtraAttrs    map[string]interface{} `json:"extra_attrs,omitempty"`
}
