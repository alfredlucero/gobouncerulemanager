package main

import (
	"fmt"
	"log"

	"gobrm/throughputrule"

	"github.com/kelseyhightower/envconfig"
)

type MySQLConfig struct {
	User     string
	Password string
	Database string
	Port     int
}

type ServerConfig struct {
	Port int
}

func main() {
	var mySQLConfig MySQLConfig
	err := envconfig.Process("mysql", &mySQLConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("MySQL Config: %+v", mySQLConfig)

	var serverConfig ServerConfig
	err = envconfig.Process("server", &serverConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Server Config: %+v", serverConfig)

	address := fmt.Sprintf(":%d", serverConfig.Port)

	a := throughputrule.App{}
	a.Initialize(mySQLConfig.User, mySQLConfig.Password, mySQLConfig.Database)
	a.Run(address)
}
