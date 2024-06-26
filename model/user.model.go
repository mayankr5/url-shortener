package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string      `json:"name"`
	Email    string      `gorm:"unique;not null" json:"email"`
	Username string      `gorm:"unique;not null" json:"username"`
	Password string      `json:"password"`
	URLs     []UserURL   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Tokens   []AuthToken `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type UserURL struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OriginalURL string    `gorm:"uniqueIndex:idx_userid_originalurl" json:"original_url"`
	ShortURL    string    `gorm:"uniqueIndex" json:"short_url"`
	Visiter     int       `gorm:"default:0" json:"visiter"`
	UserID      uuid.UUID `gorm:"uniqueIndex:idx_userid_originalurl;not null" json:"user_id"`
	Validity    time.Time `gorm:"not null" json:"validity"`
}

type AuthToken struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Token  string    `gorm:"unique;not null" json:"token"`
	UserID uuid.UUID `gorm:"not null" json:"user_id"`
}

// type Visitor struct {
// 	gorm.Model
// 	ID       uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
// 	ClientIP pgtype.Inet `gorm:"type:inet;not null" json:"client_ip"`
// }
