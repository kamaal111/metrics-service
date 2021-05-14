package main

import (
	"fmt"
	"log"

	"github.com/kamaal111/metrics-service/db"
	"github.com/kamaal111/metrics-service/router"
)

const PATH = "127.0.0.1"
const PORT = "8080"

func main() {
	pgDB := db.Connect(PATH)

	router.HandleRequests(pgDB, fmt.Sprintf("%s:%s", PATH, PORT))

	err := pgDB.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("error while closing the database connection, reason: %v\n", err))
	}

	log.Println("Connection to database closed successful.")
}
