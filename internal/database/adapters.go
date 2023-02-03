package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var ProductDb *gorm.DB

var databases map[string]*gorm.DB

func inits() {
	for k, v := range config.Database {

	}
}

func initDB(cfg *Database) *gorm.DB {
	switch cfg.Adapter {
	case "mysql":
		
	}

	return nil
}

func init() {
	var err error
	ProductDb, err = NewDriverMysql(ConfigDb)
	if err != nil {
		fmt.Println("Error in connection db News -> " + err.Error())
	}

}

type DriverMysql struct {
	db *gorm.DB
}

func NewDriverMysql(config DatabaseConfig) (*gorm.DB, error) {
	dbConn, err := connection(config)
	if err != nil {
		panic("connection database failed")
	}

	err = MigrateSchema(dbConn)
	if err != nil {
		fmt.Println("Error MigrateSchema", err)
	}

	return dbConn, nil
}

func connection(config DatabaseConfig) (*gorm.DB, error) {
	var (
		dbConn *gorm.DB
		err    error
	)

	user := config.Product.Mysql.Username
	password := config.Product.Mysql.Password
	host := config.Product.Mysql.Host
	port := config.Product.Mysql.Port
	dbname := config.Product.Mysql.Dbname

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

	currentWaitTime := 2
	trialCount := 0

	for dbConn == nil && trialCount < 5 {
		trialCount++
		dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil {
			fmt.Println("unable connecting to DB.")
			if trialCount == 5 {
				return nil, err
			}
			fmt.Println("retrying in", currentWaitTime, "seconds...")
			time.Sleep(time.Duration(currentWaitTime) * time.Second)
			currentWaitTime = currentWaitTime * 2
			dbConn = nil
		}
	}
	conn, err := dbConn.DB()
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(7)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(1 * time.Hour)

	return dbConn, err
}

var tables = []interface{}{
	//model
}

func MigrateSchema(db *gorm.DB) error {
	return db.AutoMigrate(tables...)
}
