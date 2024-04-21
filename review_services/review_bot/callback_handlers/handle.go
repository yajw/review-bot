package callback_handlers

import (
	"net/url"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yajw/review-bot/libs/logger"
)

type handler func(bot *botapi.BotAPI, msg *botapi.CallbackQuery, url *url.URL) error

var handlers = map[string]handler{
	"/submit_review": SubmitReview,
}

func Handle(bot *botapi.BotAPI, msg *botapi.CallbackQuery) error {
	u, err := url.Parse(msg.Data)
	if err != nil {
		logger.Error("url parse err:%v", err)
		return err
	}

	println(u.Path)

	if handler, ok := handlers[u.Path]; ok {
		handler(bot, msg, u)
	}
	return nil
}
