package mysql

import (
	"os"

	"github.com/jmoiron/sqlx"
)

var (
	_targetDatabase   = os.Getenv("TARGET_DATABASE")
	_connectionString = os.Getenv("MYSQL_CONNECTION_STRING")
) 

type (
MySQLxDriver struct {
	ConnectionString string
	DataSource       string
	db               *sqlx.DB
}
  IMySQLXDriver interface {
	DB() *sqlx.DB
}
)

func (m *MySQLxDriver) Init() error{
	m.ConnectionString = _connectionString
	m.DataSource = _targetDatabase
	if err := m.establishConnection(); err != nil {
		return err
	}
	return nil
}

func NewMysqlDriver() (IMySQLXDriver, error) {
	result := &MySQLxDriver{}
	if err := result.Init(); err != nil{
		return nil,err
	}
	return result,nil
}

func (m *MySQLxDriver) DB() *sqlx.DB {
	return m.db
}

func (m MySQLxDriver) establishConnection() error {
	var err error
	m.db, err = sqlx.Open(m.ConnectionString, m.DataSource)
	if err != nil {
		return err
	}
	return nil
}
