package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"

	"git.d.foundation/datcom/backend/src/app"
)

func main() {
	listenPort := os.Getenv("PORT")
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

	if listenPort == "" {
		listenPort = "8000"
		log.Println("default listening port is set to " + listenPort)
	}

	app := &app.App{}
	app, err := app.NewApp(dbConf)
	if err != nil {
		log.Fatal(err)
	}
	app.RunServer("0.0.0.0:" + listenPort)
}
