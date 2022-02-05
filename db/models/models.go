package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Mail     string `gorm:"unique;not null"`
	Password string
	Balance  float64
}

type Message struct {
	gorm.Model
	Message       string
	SenderID      uint
	Sender        Account `gorm:"foreignKey:SenderID"`
	TransactionID uint
	Transaction   Transaction `gorm:"foreignKey:TransactionID"`
}

type Transaction struct {
	gorm.Model
	Price  float64
	Status string

	PosterID uint
	Poster   Account `gorm:"foreignKey:PosterID"`
	BuyerID  uint
	Buyer    Account `gorm:"foreignKey:BuyerID"`
	AdID     uint
	Ad       Ad `gorm:"foreignKey:AdID"`
}

type Ad struct {
	gorm.Model
	Title       string
	Description string
	Price       float64
	Picture     string
	Available   bool

	PosterID uint
	Poster   Account `gorm:"foreignKey:PosterID"`
}
