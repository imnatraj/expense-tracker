package migrate

import (
	"log"

	conf "github.com/natraj/expense-tracker/config"
	"github.com/natraj/expense-tracker/models"
	"gorm.io/gorm"
)

// migrate struct in db
func migrateTable(db *gorm.DB) {
	if err := db.AutoMigrate(
		models.Hello{},
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
