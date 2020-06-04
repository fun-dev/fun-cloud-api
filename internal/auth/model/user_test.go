package model

import (
	"github.com/fun-dev/fun-cloud-api/pkg/cloudstore"
	"log"
	"testing"
)

func TestUser_GetByAccessToken(t *testing.T) {
	driver, _ := cloudstore.NewMysqlDriver()
	user := NewUserWithMySQLDriver(driver)
	// TODO: prepare test data
	// 1. deploy mysql on docker-compose
	accessToken := "test"
	result, err := user.GetByAccessToken(accessToken)
	if err != nil {
		log.Fatal(err)
	}
	expect := "test"
	if result.GoogleName != expect {
		t.Error("\n実際： ", result, "\n理想： ", expect)
	}
}
