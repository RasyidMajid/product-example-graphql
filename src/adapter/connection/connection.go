package connection

import (
	"fmt"
	"gorm.io/gorm"
	"product-service-graphql/config"
	"product-service-graphql/src/adapter/connection/mysql"
)

var Product *gorm.DB

func init() {
	var err error
	cfg := config.GetConfig()

	Product, err = mysql.NewDriverMysql(cfg.Database.Product.Mysql)
	if err != nil {
		if err != nil {
			fmt.Println("Error in connection db News -> " + err.Error())
		}
	}
}
