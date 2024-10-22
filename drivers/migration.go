package drivers

import (
	"catering-api/models/model"
	"gorm.io/gorm"
)

func DBMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.Employee{},
	)

	if err != nil {
		return err
	}

	return nil
}
