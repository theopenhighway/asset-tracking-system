package utils

import (
	"asset-tracking-system/dto"
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, msg string, data any, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(dto.ResponseMessage{
		Msg:  msg,
		Data: data,
	})
}
