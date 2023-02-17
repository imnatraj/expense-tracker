package helper

import (
	"encoding/json"
	"net/http"
)

type response struct {
	http.ResponseWriter
}

func Res(w http.ResponseWriter) *response {
	return &response{w}
}

type successResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (w *response) Success(status int, data interface{}) {
	payload := &successResponse{
		Status: status,
		Data:   data,
	}
	bin, err := json.Marshal(payload)
	if err != nil {
		w.Error(http.StatusInternalServerError, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bin)
}

type errorResponse struct {
	Status int `json:"status"`
	Error  struct {
		Message string `json:"message"`
	} `json:"error"`
}

func (w *response) Error(status int, err error) {
	payload := &errorResponse{
		Status: status,
		Error: struct{Message string "json:\"message\""}{
			Message: err.Error(),
		},
	}
	bin, err := json.Marshal(payload)
	if err != nil {
		w.Error(http.StatusInternalServerError, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bin)
}
