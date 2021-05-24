package main

import (
	"fmt"
	"gobrm/bouncerule"
	"os"
)

func main() {
	a := bouncerule.App{}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("SERVER_PORT")
	address := fmt.Sprintf(":%s", port)

	fmt.Printf("Running server on port %s...", port)

	a.Initialize(user, password, dbname)
	a.Run(address)
}
