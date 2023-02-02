package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	dbConfig "product-service-graphql/config/db"
	"product-service-graphql/src/adapter/presenter/model"
	"time"
)

type DriverMysql struct {
	db *gorm.DB
}

func NewDriverMysql(config dbConfig.Database) (*gorm.DB, error) {
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

func connection(config dbConfig.Database) (*gorm.DB, error) {
	var (
		dbConn *gorm.DB
		err    error
	)

	user := config.Username
	password := config.Password
	host := config.Host
	port := config.Port
	dbname := config.Dbname

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
	model.ProductModel{},
}

func MigrateSchema(db *gorm.DB) error {
	return db.AutoMigrate(tables...)
}
