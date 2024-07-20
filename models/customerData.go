package models

import (
	"gorm.io/gorm"
)

type CustomerData struct {
	gorm.Model
	CustomerId  *uint         `json:"customerId"`
	Platform    *string       `json:"platform"`
	Account     *string       `json:"account"`
	Transaction []Transaction `gorm:"foreignKey:CustomerDataId" json:"transaction"`
}

func MigrateCustomerData(db *gorm.DB) error {
	err := db.AutoMigrate(&CustomerData{})
	return err
}
