package common

import (
	"net/http"

	"github.com/yajw/review-bot/libs/logger"
)

func HandleError(err error, w http.ResponseWriter) {
	logger.Error("error: %+v", err)
	w.WriteHeader(http.StatusInternalServerError)
}
