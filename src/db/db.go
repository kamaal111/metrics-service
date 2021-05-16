package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"github.com/kamaal111/metrics-service/src/models"
)

func Connect(dbPath string) *pg.DB {
	options := &pg.Options{
		User:     "postgres",
		Password: "pass",
		Addr:     fmt.Sprintf("%s:5432", dbPath),
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
	err := createAppsTable(pgDB)
	if err != nil {
		return err
	}
	err = createMetricsTable(pgDB)
	if err != nil {
		return err
	}
	return nil
}

func createAppsTable(pgDB *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := pgDB.Model((*models.Apps)(nil)).CreateTable(options)
	return err
}

func createMetricsTable(pgDB *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := pgDB.Model((*models.MetricsTable)(nil)).CreateTable(options)
	return err
}
