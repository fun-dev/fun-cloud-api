package mysql

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var (
	_targetDatabase   = os.Getenv("TARGET_DATABASE")
	_connectionString = os.Getenv("MYSQL_CONNECTION_STRING")
) 

type MySQLxDriver struct {
	ConnectionString string
	DataSource       string
	DB               *sqlx.DB
}

type IMySQLXDriver interface {
	Database() *sqlx.DB
}

func (m *MySQLxDriver) Init() {
	m.ConnectionString = _connectionString
	m.DataSource = _targetDatabase
	if err := m.establishConnection(); err != nil {
		log.Fatal(err)
	}
}

func NewMysqlDriver() IMySQLXDriver {
	result := &MySQLxDriver{}
	result.Init()
	return result
}

func (m *MySQLxDriver) Database() *sqlx.DB {
	return m.DB
}

func (m MySQLxDriver) establishConnection() error {
	var err error
	m.DB, err = sqlx.Open(m.ConnectionString, m.DataSource)
	if err != nil {
		return err
	}
	return nil
}
