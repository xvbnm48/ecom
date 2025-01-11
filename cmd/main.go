package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/xvbnm48/ecom/cmd/api"
	"github.com/xvbnm48/ecom/config"
	"github.com/xvbnm48/ecom/db"
)

func main() {
	db, err := db.NewMySqlStorage(mysql.Config{
		User:                 config.Evns.DBUser,
		Passwd:               config.Evns.DBPassword,
		Addr:                 config.Evns.DBAddress,
		DBName:               config.Evns.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	fmt.Print("db address:", config.Evns.DBAddress)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)
	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if ping := db.Ping(); ping != nil {
		log.Fatal(ping)
	}
	log.Println("Database connected")
}
