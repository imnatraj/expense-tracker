package hello

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/natraj/expense-tracker/helper"
	"github.com/natraj/expense-tracker/models"
	"github.com/natraj/expense-tracker/service"
)

// api endpoint for hello will be handled here

func Route(ctx *models.HandlerCtx) func(c chi.Router) {
	return func(r chi.Router) {
		r.Post("/", create(ctx))
		r.Get("/{id}", getOne(ctx))
		r.Put("/{id}", updateOne(ctx))
		r.Get("/", getMany(ctx))
	}
}

func create(ctx *models.HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metaData, err := helper.Req(r).MetaData()
		if err != nil {
			helper.Res(w).Error(http.StatusBadRequest, err)
			return
		}

		fmt.Println(metaData)

		if metaData.Mode == models.MetaDataModeWeb {
			fmt.Println("request from web")
		}
		if metaData.Mode == models.MetaDataModeApp {
			fmt.Println("request from app")
		}

		var body createHelloBody

		err = helper.Req(r).B(&body)
		if err != nil {
			helper.Res(w).Error(http.StatusBadRequest, err)
			return
		}

		payload := &models.Hello{
			Name: body.Name,
		}

		err = service.Hello(ctx).Create(payload)
		if err != nil {
			helper.Res(w).Error(http.StatusInternalServerError, err)
			return
		}

		helper.Res(w).Success(http.StatusCreated, body)
	}
}

func getOne(ctx *models.HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pId := helper.Req(r).P("id")
		id, err := helper.StringToUint(pId)
		if err != nil {
			helper.Res(w).Error(http.StatusBadRequest, err)
			return
		}
		data, err := service.Hello(ctx).GetOne(id)
		if err != nil {
			helper.Res(w).Error(http.StatusInternalServerError, err)
			return
		}

		helper.Res(w).Success(http.StatusCreated, data)
	}
}

func updateOne(ctx *models.HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pId := helper.Req(r).P("id")
		id, err := helper.StringToUint(pId)
		if err != nil {
			helper.Res(w).Error(http.StatusBadRequest, err)
			return
		}

		var body createHelloBody

		err = helper.Req(r).B(&body)
		if err != nil {
			helper.Res(w).Error(http.StatusBadRequest, err)
			return
		}

		payload := &models.Hello{
			Name: body.Name,
			ID:   id,
		}

		err = service.Hello(ctx).Update(id, payload)
		if err != nil {
			helper.Res(w).Error(http.StatusInternalServerError, err)
			return
		}

		helper.Res(w).Success(http.StatusCreated, body)
	}
}

func getMany(ctx *models.HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hellos, err := service.Hello(ctx).GetMany()
		if err != nil {
			helper.Res(w).Error(http.StatusInternalServerError, err)
			return
		}

		helper.Res(w).Success(http.StatusCreated, hellos)
	}
}
