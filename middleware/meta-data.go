package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/natraj/expense-tracker/helper"
	"github.com/natraj/expense-tracker/models"
)

func MetaData() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mode := r.Header.Get("X-Mode")
			channel := r.Header.Get("X-Channel")
			validMode := models.MetaDataMode(mode)
			validChanel := models.MetaDataChannel(channel)

			if !(len(mode) != 0 && len(channel) != 0) {
				helper.Res(w).Error(http.StatusBadRequest, errors.New("invalid meta data"))
				return
			}

			metaData := &models.MetaData{
				Mode:    validMode,
				Channel: validChanel,
			}
			ctx := context.WithValue(
				r.Context(),
				models.RequestDataMetaData,
				metaData,
			)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
