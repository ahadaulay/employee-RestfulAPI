package helpers

import (
	"catering-api/drivers"

	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {

	configDatabase := drivers.ConfigDB{
		DB_USERNAME: GetConfig("DB_USERNAME"),
		DB_PASSWORD: GetConfig("DB_PASSWORD"),
		DB_HOST:     GetConfig("DB_HOST"),
		DB_PORT:     GetConfig("DB_PORT"),
		DB_NAME:     GetConfig("DB_NAME"),
	}

	db := configDatabase.InitDB()

	err := drivers.DBMigration(db)

	if err != nil {
		panic(err)
	}

	return db
}
