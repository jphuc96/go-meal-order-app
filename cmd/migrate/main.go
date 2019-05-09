package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"git.d.foundation/datcom/backend/src/store"
)

func main() {
	Host := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	SSLMode := os.Getenv("DB_SSL")

	if Host == "" || Port == "" || User == "" || Password == "" || DBName == "" || SSLMode == "" {
		log.Fatal("missing environments")
	}

	connString := "host=" + Host +
		" port=" + Port +
		" user=" + User +
		" password=" + Password +
		" dbname=" + DBName +
		" sslmode=" + SSLMode

	store, err := store.NewPostgresMigrator(connString)
	if err != nil {
		fmt.Println(err)
	}
	_ = store

	flag := os.Args[1]
	if flag != "up" && flag != "down" {
		log.Fatalln("up down")
	}

	switch flag {
	case "up":
		err = store.MigrateDB()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "down":
		err = store.ReverseDB()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
