package service

import (
	"time"

	"github.com/yajw/review-bot/review_services/common"
)

var (
	mockedTemplateStore = map[string]*common.ReviewTemplate{
		"transaction_list": {
			SceneKey:   "transaction_list",
			Template:   "Hi %s, how is our logistic service?",
			CreateTime: time.Now(),
		},
		"chat.order": {
			SceneKey:   "chat.order",
			Template:   "%s rate our logistic service here",
			CreateTime: time.Now(),
		},
		"default": {
			SceneKey:   "default",
			Template:   "Please give us some feedback",
			CreateTime: time.Now(),
		},
	}
)

func CreateReviewTemplate(template *common.ReviewTemplate) {
	mockedTemplateStore[template.SceneKey] = template
}

func GetReviewTemplateBySceneKey(key string) *common.ReviewTemplate {
	if template, ok := mockedTemplateStore[key]; ok {
		return template
	}

	return mockedTemplateStore["default"]
}
