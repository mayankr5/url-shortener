package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Username string    `gorm:"unique;not null" json:"username"`
	Password string    `validate:"min=8" json:"password"`
}

type UserURLs struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OriginalURL string    `gorm:"unique" json:"original_url"`
	ShortURL    string    `gorm:"unique" json:"short_url"`
	Visiter     int       `gorm:"default:0" json:"visiter"`
	UserID      uuid.UUID `gorm:"not null" json:"user_id"`
}

type AuthToken struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Token  string    `gorm:"unique;not null" json:"token"`
	UserID uuid.UUID `json:"user_id"`
}
