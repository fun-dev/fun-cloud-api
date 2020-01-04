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
		BaseModel
		IconUrl     string `json:"icon_url" db:"icon_url"`
		GoogleName  string `json:"google_name" db:"google_name"`
		AccessToken string `json:"-" db:"access_token, primarykey"`
		// --- other struct ---
		MySQLDriver mysql.IMySQLXDriver
	}
)

// Constructor with DB Connection
func NewUserWithMySQLDriver(mysqlDriver mysql.IMySQLXDriver) IUser {
	result := &User{}
	result.MySQLDriver = mysqlDriver
	return result
}

// Normal Constructor
func NewUser(iconUrl, googleName, accessToken string) *User {
	return &User{
		IconUrl:     iconUrl,
		GoogleName:  googleName,
		AccessToken: accessToken,
	}
}

func (u *User) GetByAccessToken(accessToken string) (*User, error) {
	result := &User{}
	if err := u.MySQLDriver.DB().Get(&result, "SELECT * FROM users WHERE access_token=$1", accessToken); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *User) Create(item User) error {
	// insert user model to db by sqlx
	_, err := u.MySQLDriver.DB().NamedExec("INSERT INTO users (icon_url, google_name, access_token) VALUES (:icon_url, :google_name, :access_token)", &item)
	if err != nil {
		//TODO: implement error handling
		return err
	}
	return nil
}