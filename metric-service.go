package main

import (
	"fmt"
	"log"

	"github.com/kamaal111/metrics-service/db"
	"github.com/kamaal111/metrics-service/router"
)

func main() {
	pgDB := db.Connect()

	router.HandleRequests(pgDB)

	err := pgDB.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("error while closing the database connection, reason: %v\n", err))
	}

	log.Println("Connection to database closed successful.")
}
