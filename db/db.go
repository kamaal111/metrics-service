package db

import (
	"errors"
	"log"

	"github.com/go-pg/pg/v10"
)

func Connect() *pg.DB {
	options := &pg.Options{
		User:     "postgres",
		Password: "pass",
		Addr:     "127.0.0.1:5432",
	}

	pgDB := pg.Connect(options)
	if pgDB == nil {
		log.Fatal(errors.New("failed to connect to database"))
	}

	log.Println("Connection to database successful.")

	err := createSchema(pgDB)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully created database schemas")

	return pgDB
}

func createSchema(pgDB *pg.DB) error {
	return nil
}
