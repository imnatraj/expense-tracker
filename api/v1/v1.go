package v1

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/natraj/expense-tracker/api/v1/hello"
	conf "github.com/natraj/expense-tracker/config"
	"github.com/natraj/expense-tracker/models"
)

func ApiV1(env string) func(c chi.Router) {

	envFileName := ".env." + env
	var config models.Config
	conf.GetEnv(envFileName, "yml", ".", &config)

	fmt.Println(config.App.Server.Port)

	var ctx models.HandlerCtx

	userSession := conf.GetRedisDbClient(config.App.Redis.Host, config.App.Redis.Password, conf.RedisDBTypeSession)
	otpSession := conf.GetRedisDbClient(config.App.Redis.Host, config.App.Redis.Password, conf.RedisDBTypeOTP)
	pConf := config.App.Postgres
	p2Conf := config.App.Postgres2
	postgresDB := conf.NewPostgresDB(pConf.Host, pConf.DBName, pConf.Password, pConf.User, pConf.Port)
	postgres2DB := conf.NewPostgresDB(p2Conf.Host, p2Conf.DBName, p2Conf.Password, p2Conf.User, p2Conf.Port)
	ctx.UserSessionDB = userSession
	ctx.OtpSessionDB = otpSession
	ctx.DB = postgresDB
	ctx.AnotherDB = postgres2DB

	return func(r chi.Router) {
		r.Route("/hello", hello.Route(&ctx))
	}
}
