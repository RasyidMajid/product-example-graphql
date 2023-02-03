package database

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os"
)

type Config struct {
	Database map[string]*Database
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Adapter  string
}

func (d *Database) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Dbname,
	)
}

var config Config

func init() {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configor.Load(&config, dir+"shared/database.yaml")
}
