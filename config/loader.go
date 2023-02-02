package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"product-service-graphql/config/db"
	"product-service-graphql/config/server"
	"strings"
)

type config struct {
	Server   server.ServerList
	Database db.DatabaseList
}

var cfg config

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(dir + "/config/server")
	viper.SetConfigType("yaml")
	viper.SetConfigName("server.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load Server config: %v", err))
	}

	viper.AddConfigPath(dir + "/config/db")
	viper.SetConfigType("yaml")
	viper.SetConfigName("database.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load Database config: %v", err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}

	viper.Unmarshal(&cfg)

	fmt.Println("=============================")
	dataByte, _ := json.Marshal(cfg)
	fmt.Println(string(dataByte))
	fmt.Println("=============================")
}
func GetConfig() *config {
	return &cfg
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(env) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}
