package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kamaal111/metrics-service/src/db"
	"github.com/kamaal111/metrics-service/src/router"
)

func main() {
	PATH := os.Getenv("APP_PATH")
	if len(PATH) < 1 {
		PATH = "127.0.0.1"
	}
	PORT := os.Getenv("APP_PORT")
	if len(PORT) < 1 {
		PORT = "8080"
	}

	pgDB := db.Connect(PATH)

	router.HandleRequests(pgDB, fmt.Sprintf("%s:%s", PATH, PORT))

	err := pgDB.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("error while closing the database connection, reason: %v\n", err))
	}

	log.Println("Connection to database closed successful.")
}
