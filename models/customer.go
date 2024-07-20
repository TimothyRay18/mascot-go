package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	// Id               uint           `gorm:"primary key;autoIncrement" json:"id"`
	gorm.Model
	Name         *string        `json:"name"`
	Cis          *string        `json:"cis"`
	BcaId        *string        `json:"bcaId"`
	CustomerData []CustomerData `gorm:"foreignkey:CustomerId" json:"customerData"`
	Transaction  []Transaction  `gorm:"foreignKey:CustomerId" json:"transaction"`
}

func MigrateCustomer(db *gorm.DB) error {
	err := db.AutoMigrate(&Customer{})
	return err
}
