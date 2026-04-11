// models/index.go
package models

import (
	"log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	log.Println("⏳ Migrating database tables...")

	err := db.AutoMigrate(
		&Member{},
		// &OtherModel{},
	)

	if err != nil {
		return err
	}

	log.Println("✅ All tables migrated successfully!")
	return nil
}
