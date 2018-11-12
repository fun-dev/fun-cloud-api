package dbmodels

import "time"

// User is syncronizing database model
type User struct {
	Id          int64
	IconURL     string
	GoogleName  string
	AccessToken string
	Created     time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
}
