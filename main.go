package main

import (
	"log"

	"github.com/themelancholyspirit/airline-reservation-system/api"
	"github.com/themelancholyspirit/airline-reservation-system/database"
	"github.com/themelancholyspirit/airline-reservation-system/storage"
)

func main() {

	dbConfig := database.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "mysecretpassword",
		DBName:   "postgres",
		SSLMode:  "disable",
	}

	db, err := database.NewPostgreDB(dbConfig)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	storage := storage.NewPostgreStorage(db)

	server := api.NewServer(":8080", storage)
	server.SetupRoutes()
	server.Router.Run()
}
