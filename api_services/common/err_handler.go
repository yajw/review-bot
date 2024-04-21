package common

import (
	"encoding/json"
	"net/http"

	"github.com/yajw/review-bot/libs/logger"
)

func HandleError(err error, w http.ResponseWriter) {
	logger.Error("error: %+v", err)
	w.WriteHeader(http.StatusInternalServerError)
}

func JsonResp(w http.ResponseWriter, o interface{}) {
	w.Header().Set("Content-Type", "application/json")
	bts, err := json.Marshal(o)
	if err != nil {
		HandleError(err, w)
		return
	}

	w.Write(bts)
}
