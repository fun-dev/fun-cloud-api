package config

import (
	"fmt"
	"testing"
)

func TestGetSQLUser(t *testing.T) {
	fmt.Println(GetSQLUser())
}

func TestGetSQLPass(t *testing.T) {
	fmt.Println(GetSQLPass())
}
func TestGetSQLHost(t *testing.T) {
	fmt.Println(GetSQLHost())
}
func TestGetSQLPort(t *testing.T) {
	fmt.Println(GetSQLPort())
}

func TestGetSQLDB(t *testing.T) {
	fmt.Println(GetSQLDB())
}

func TestGetKubeConfigPath(t *testing.T) {
	fmt.Println(GetKubeConfigPath())
}
