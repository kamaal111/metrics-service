package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"github.com/kamaal111/metrics-service/src/models"
)

var PGDatabase *pg.DB

func Connect() {
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	if POSTGRES_USER == "" {
		log.Fatalln("POSTGRES_USER is undefined")
	}
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	if POSTGRES_PASSWORD == "" {
		log.Fatalln("POSTGRES_PASSWORD is undefined")
	}
	DP_PORT := os.Getenv("DP_PORT")
	if DP_PORT == "" {
		log.Fatalln("DP_PORT is undefined")
	}
	DB_PATH := os.Getenv("DB_PATH")

	var err error
	var options *pg.Options

	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL != "" {
		parsedDBUrl, err := pg.ParseURL(DATABASE_URL)
		if err != nil {
			log.Fatalln(err)
		}
		options = parsedDBUrl
	} else {
		options = &pg.Options{
			User:     POSTGRES_USER,
			Password: POSTGRES_PASSWORD,
			Addr:     fmt.Sprintf("%s:%s", DB_PATH, DP_PORT),
		}
	}

	PGDatabase = pg.Connect(options)
	if PGDatabase == nil {
		log.Fatal(errors.New("failed to connect to database"))
	}

	PGDatabase.AddQueryHook(dbLogger{})

	log.Println("Connection to database successful.")

	err = createSchema(PGDatabase)
	if err != nil {
		log.Fatal(err)
	}
}

func BulkSaveMetrics(pgDB *pg.DB, metrics []models.MetricsTable) error {
	_, err := pgDB.Model(&metrics).Insert()
	return err
}

func GetAppByBundleIdentifier(pgDB *pg.DB, bundleIdentifier string) (models.AppsTable, error) {
	var app models.AppsTable
	err := pgDB.Model(&app).Where("bundle_identifier = ?", bundleIdentifier).Select()
	return app, err
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
	err := pgDB.Model((*models.AppsTable)(nil)).CreateTable(options)
	return err
}

func createMetricsTable(pgDB *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := pgDB.Model((*models.MetricsTable)(nil)).CreateTable(options)
	return err
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, err := q.UnformattedQuery()
	if err != nil {
		log.Printf("error %s\n", err.Error())
		return nil
	}
	log.Println(string(query))
	return nil
}
