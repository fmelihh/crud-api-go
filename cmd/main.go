package main

import (
	"log"

	"github.com/fmelihh/crud-api-go/cmd/api"
	"github.com/fmelihh/crud-api-go/config"
	"github.com/fmelihh/crud-api-go/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser",
		Passwd:               "asd",
		Addr:                 "127.0.1:3306",
		DBName:               "ecom",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
