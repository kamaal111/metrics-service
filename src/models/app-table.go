package models

import (
	"github.com/go-pg/pg/v10"
)

type AppsTable struct {
	tableName        struct{} `pg:"apps"`
	ID               int      `pg:"id,pk"`
	BundleIdentifier string   `pg:"bundle_identifier,unique"`
	AccessToken      string   `pg:"access_token"`
}

func (app *AppsTable) Save(pgDB *pg.DB) error {
	_, err := pgDB.Model(app).Insert()
	return err
}

func (app *AppsTable) GetMetrics(pgDB *pg.DB) ([]MetricsTable, error) {
	var metrics []MetricsTable
	err := pgDB.Model(&metrics).Where("app_id = ?", app.ID).Select()
	return metrics, err
}
