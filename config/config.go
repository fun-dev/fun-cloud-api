package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	SQLUser    string `envconfig:"SQL_USER" default:"root"`
	SQLPass    string `envconfig:"SQL_PASS" default:"zJ6pF57r54JH"`
	SQLHost    string `envconfig:"SQL_HOST" default:"127.0.0.1"`
	SQLPort    string `envconfig:"SQL_PORT" default:"3306"`
	SQLDB      string `envconfig:"SQL_DB" default:"prac"`
	BucketName string `envconfig:"BUCKET_NAME" default:"gin-practice"`
}

const (
	prefix = "APP"
)

var (
	c Config
)

func init() {
	envconfig.MustProcess(prefix, &c)
}

func GetSQLUser() string {
	return c.SQLUser
}

func GetSQLPass() string {
	return c.SQLPass
}
func GetSQLHost() string {
	return c.SQLHost
}
func GetSQLPort() string {
	return c.SQLPort
}

func GetSQLDB() string {
	return c.SQLDB
}

func GetBucketName() string {
	return c.BucketName
}
