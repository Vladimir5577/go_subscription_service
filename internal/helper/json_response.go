package helper

import (
	"encoding/json"
	"net/http"
)

type JsonResponseType struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func JsonResponse(w http.ResponseWriter, data any, statusCode int) {
	resp := JsonResponseType{}
	if statusCode == http.StatusOK ||
		statusCode == http.StatusCreated {
		resp.Success = true
	} else {
		resp.Success = false
	}

	resp.Data = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
