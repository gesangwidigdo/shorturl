package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&Url{}); err != nil {
		return err
	}

	return nil
}
