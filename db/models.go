package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(50);unique;not null"`
	Password string `gorm:"type:varchar(50);not null"`
}

type Deck struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);not null;uniqueIndex:ix_deck_name_userid" yaml:"name"`
	UserID uint   `gorm:"not null;uniqueIndex:ix_deck_name_userid"`
	User   User   `gorm:"constraint:OnDelete:CASCADE"`
}

type Card struct {
	gorm.Model
	Front  string `gorm:"type:text;not null" yaml:"front"`
	Back   string `gorm:"type:text;not null" yaml:"back"`
	DeckID uint   `gorm:"not null"`
	Deck   Deck   `gorm:"constraint:OnDelete:CASCADE"`
}
