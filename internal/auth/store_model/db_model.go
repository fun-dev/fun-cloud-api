package store_model

import "time"

type BaseModel struct {
	ID        uint       `db:"id" json:"-"`
	CreatedAt *time.Time `db:"created_at" json:"-"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
}
