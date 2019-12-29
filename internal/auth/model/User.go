package model

import (
	"github.com/fun-dev/fun-cloud-api/pkg/mysql"
)

type (
	IUser interface {
		//TODO: implement User.Create by MySQLDriver
		Create(item User) error
		GetByAccessToken(accessToken string) (*User, error)
	}
	User struct {
		IconUrl      string `json:"icon_url" db:"icon_url"`
		GoogleName   string `json:"google_name" db:"google_name"`
		AccesesToken string `json:"-" db:"accesstoken, primarykey"`
		// --- other struct ---
		MySQLDriver mysql.IMySQLXDriver
	}
)

func NewUserWithMySQLDriver(mysqlDriver mysql.IMySQLXDriver) IUser {
	result := &User{}
	result.MySQLDriver = mysqlDriver
	return result
}


func (u *User) GetByAccessToken(accessToken string) (*User, error) {
	result := &User{}

	if err := u.MySQLDriver.DB().Select(&result, "SELECT * FROM"); err != nil {
		return nil, err
	}
	return result, nil
}

func NewUser() IUser {
	result := &User{}
	//result.Init()
	return result
}

//func NewUser(User.IconUrl,User.GoogleName,User.AccesesToken){
//
//}

//func NewUser(iconURL, googleName, accessToken) IUser {
//	result := &User{}
//	result.IconUrl = iconURL
//	result.GoogleName = googleName
//	result.AccesesToken = accessToken
//	return result
//}
