package command_handlers

import (
	"fmt"
	"net/url"
	"strconv"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yajw/review-bot/libs/logger"
	"github.com/yajw/review-bot/review_services/customer_review_service"
)

func buildReviewKeyBoard(sid int64) botapi.InlineKeyboardMarkup {
	builder := func(content string) string {
		return buildSchema("/submit_review", map[string]interface{}{"c": content, "sid": sid})
	}

	var numericKeyboard = botapi.NewInlineKeyboardMarkup(
		botapi.NewInlineKeyboardRow(
			botapi.NewInlineKeyboardButtonData("Good", builder("good")),
			botapi.NewInlineKeyboardButtonData("Average", builder("average")),
			botapi.NewInlineKeyboardButtonData("Bad", builder("bad")),
		),
	)
	return numericKeyboard
}

func buildSchema(path string, args map[string]interface{}) string {
	u, _ := url.Parse(path)
	q := u.Query()
	for k, v := range args {
		q.Set(k, fmt.Sprintf("%v", v))
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func MockReceiveProduct(bot *botapi.BotAPI, msg *botapi.Message) error {
	sceneID, err := strconv.ParseInt(string(msg.CommandArguments()[0]), 10, 64)
	if err != nil {
		logger.Error("parse sid err:%v", err)
		return err
	}
	return ReceiveProduct(bot, msg.Chat.ID, msg.Chat.UserName, "chat.order", int64(sceneID))
}

func ReceiveProduct(bot *botapi.BotAPI, chatID int64, userName string, sceneKey string, sceneID int64) error {
	tpl, err := customer_review_service.GetReviewService().GetReviewTemplate("chat.order")
	if err != nil {
		return err
	}

	text := fmt.Sprintf(tpl.Template, userName)
	reply := botapi.NewMessage(chatID, text)
	reply.ReplyMarkup = buildReviewKeyBoard(sceneID)

	if _, err := bot.Send(reply); err != nil {
		return err
	}

	return nil
}
