package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id              uint       `gorm:"primary key;autoIncrement" json:"id"`
	CustomerId      uint       `gorm:"type:bigint" json:"customerId"`
	CustomerDataId  uint       `gorm:"type:bigint" json:"customerDataId"`
	AppPlatform     *string    `json:"appPlatform"`
	IncomingChannel *string    `json:"incomingChannel"`
	InteractionTime *time.Time `json:"interactionTime"`
	TicketID        *string    `json:"ticketId"`
}

func MigrateTransaction(db *gorm.DB) error {
	err := db.AutoMigrate(&Transaction{})
	return err
}
