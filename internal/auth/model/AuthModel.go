package model

import (
	"github.com/fun-dev/fun-cloud-api/pkg/mysql"
)


type (
	IUser interface {
		Get(accessToken string) (*User, error)
	}
	User struct {
		IconUrl      string `json:"icon_url"`  `db:"icon_url"`
		GoogleName   string `json:"google_name"`  `db:"google_name"`
		AccesesToken string `json:"-"`  `db:"accesstoken, primarykey"`
		// --- other struct ---
		MySQLDriver mysql.IMySQLXDriver
	}
)

// func NewUserWithMySQLDriver (User.IconUrl,User.GoogleName,User.AccesesToken,mysql.IMySQLXDriver)*mysql.IMySQLXDriver {
//  	return {
//  		IMySQLXDriver
//  	}
// }
 
func NewUserWithMySQLDriver(mysqlDriver mysql.IMySQLXDriver) IUser {
	result := &User{}
	result.MySQLDriver = mysqlDriver
	return result
}

func NewUser(iconURL, googleName, accessToken) IUser {
	result := &User{}
	result.IconUrl = iconURL
	result.GoogleName = googleName
	result.AccesesToken = accessToken
	return result
}

func (u *User) Get(accessToken string) (*User, error) { 
	result := &User{}
	u.mysql.IMySQLXDriver.Database().Slect(&result,"SELECT * FROM")
} 

func NewUser(User.IconUrl,User.GoogleName,User.AccesesToken){

}


