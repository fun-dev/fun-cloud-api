package dbmodels

import "time"

type Item struct {
	Id      int64
	Name    string
	UserId  int64
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
