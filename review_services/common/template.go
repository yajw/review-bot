package common

import (
	"time"
)

type ReviewTemplate struct {
	SceneKey   string    `json:"scene_key,omitempty"` // 业务自定义的场景 key
	Template   string    `json:"template,omitempty"`  // 模板
	CreateTime time.Time `json:"create_time,omitempty"`
}
