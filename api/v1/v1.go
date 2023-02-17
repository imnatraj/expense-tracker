package v1

import (
	"fmt"

	"imnatraj/expense-tracker/api/v1/user"
	conf "imnatraj/expense-tracker/config"
	"imnatraj/expense-tracker/models"

	"github.com/go-chi/chi/v5"
)

func ApiV1(env string) func(c chi.Router) {

	envFileName := ".env." + env
	var config models.Config
	conf.GetEnv(envFileName, "yml", ".", &config)

	fmt.Println(config.App.Server.Port)

	var ctx models.HandlerCtx

	pConf := config.App.Postgres
	postgresDB := conf.NewPostgresDB(pConf.Host, pConf.DBName, pConf.Password, pConf.User, pConf.Port)
	ctx.DB = postgresDB

	return func(r chi.Router) {
		r.Route("/user", user.Route(&ctx))
	}
}
