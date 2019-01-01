package dbmodels

import "time"

// User is syncronizing database model
type User struct {
	Id          int64
	IconUrl     string
	GoogleName  string
	AccessToken string
	Email       string
	Created     time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
}
