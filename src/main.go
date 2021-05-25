package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kamaal111/metrics-service/src/db"
	"github.com/kamaal111/metrics-service/src/router"
)

func main() {
	SECRET_TOKEN := os.Getenv("SECRET_TOKEN")
	if SECRET_TOKEN == "" {
		log.Fatalln("SECRET_TOKEN is undefined")
	}

	db.Connect()

	PORT := os.Getenv("APP_PORT")
	router.HandleRequests(fmt.Sprintf(":%s", PORT))

	err := db.PGDatabase.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("error while closing the database connection, reason: %v\n", err))
	}

	log.Println("Connection to database closed successful.")
}
