package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"

	"git.d.foundation/datcom/backend/src/app"
)

func main() {
	dbConf := &app.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL"),
	}
	if dbConf.Host == "" || dbConf.Port == "" || dbConf.User == "" || dbConf.Password == "" || dbConf.DBName == "" || dbConf.SSLMode == "" {
		log.Fatal("missing environments")
	}

	app := &app.App{}
	app, err := app.NewApp(dbConf)
	if err != nil {
		log.Fatal(err)
	}
	app.RunServer("127.0.0.1:8080")
}
