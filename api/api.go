package api

import (
	"encoding/json"
	"net/http"
)

// Define params for the API request
type CoinBalanceRequest struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int
	Balance int64
}

type ErrorResponse struct {
	Code    int
	Message string
}

func returnError(h http.ResponseWriter, code int, message string) {
	error := ErrorResponse{
		Code:    code,
		Message: message,
	}

	h.Header().Set("Content-Type", "application/json")
	h.WriteHeader(code)

	json.NewEncoder(h).Encode(error)
}

var (
	UserCausedError = func(h http.ResponseWriter, err error) {
		returnError(h, http.StatusBadRequest, err.Error())
	}
	ServerCausedError = func(h http.ResponseWriter) {
		returnError(h, http.StatusInternalServerError, "Internal Server Error")
	}
)
