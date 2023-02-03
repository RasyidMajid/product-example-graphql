package server

import (
	"github.com/jinzhu/configor"
	"os"
)

type Server struct {
	Graphql Servers `json:"graphql"`
}

type Servers struct {
	TLS     bool   `yaml:"tls"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Timeout int    `yaml:"timeout"`
}

var ConfigServer Server

func init() {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	loc := dir + "/shared/server.yaml"
	err = configor.Load(&ConfigServer, loc)
	if err != nil {

	}
}
