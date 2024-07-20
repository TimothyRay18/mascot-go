package models

import "gorm.io/gorm"

type UserApp struct {
	gorm.Model
	Username   *string `json:"username"`
	Password   *string `json:"password"`
	ViewAccess *string `json:"viewAccess"`
	Level      *int    `json:"level"`
}

func MigrateUserApp(db *gorm.DB) error {
	err := db.AutoMigrate(&UserApp{})
	return err
}
