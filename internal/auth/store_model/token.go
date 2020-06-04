package store_model

import "time"

type FCPAccessToken struct {
	ID           string     `gorm:"primary_key"`
	RefreshToken string     `db:"refresh_token" gorm:"not null"`
	ExpiredAt    *time.Time `db:"expired_at"`
	CreatedAt    *time.Time `db:"created_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}
