package dbmodels

import "time"

type User struct {
	Id      int64
	Name    string
	Age     int
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
