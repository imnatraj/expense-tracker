package migrate

import (
	"log"

	conf "imnatraj/expense-tracker/config"
	"imnatraj/expense-tracker/models"

	"gorm.io/gorm"
)

// migrate struct in db
func migrateTable(db *gorm.DB) {
	if err := db.AutoMigrate(
		models.User{},
	); err != nil {
		log.Fatal(err)
	}
}

func RunMigration(env string) {
	envFileName := ".env." + env
	var config models.Config
	conf.GetEnv(envFileName, "yml", ".", &config)

	pConf := config.App.Postgres
	db := conf.NewPostgresDB(pConf.Host, pConf.DBName, pConf.Password, pConf.User, pConf.Port)
	migrateTable(db)
}
