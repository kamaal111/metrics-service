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
		// TODO: Get this from ENV
		User: "postgres",
		// TODO: Get this from ENV
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

func GetOrCreateAppByBundleIdentifier(pgDB *pg.DB, bundleIdentifier string) (models.AppsTable, error) {
	app, err := getAppByBundleIdentifier(pgDB, bundleIdentifier)
	if err == pg.ErrNoRows {
		app = models.AppsTable{
			BundleIdentifier: bundleIdentifier,
		}
		err = app.Save(pgDB)
		return app, err
	}
	return app, err
}

func getAppByBundleIdentifier(pgDB *pg.DB, bundleIdentifier string) (models.AppsTable, error) {
	var app models.AppsTable
	err := pgDB.Model(&app).Where("bundle_identifier = ?", bundleIdentifier).Select()
	return app, err
}

func GetAppWithMetricsByBundleIdentifier(pgDB *pg.DB, bundleIdentifier string) (models.AppsTable, error) {
	var app models.AppsTable
	err := pgDB.Model(&app).Where("bundle_identifier = ?", bundleIdentifier).Relation("Metrics").Select()
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
