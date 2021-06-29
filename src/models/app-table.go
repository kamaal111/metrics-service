package models

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
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

func (app *AppsTable) GetMetrics(pgDB *pg.DB, queries map[string]string) ([]MetricsTable, error) {
	var metrics []MetricsTable
	startingQuery := pgDB.Model(&metrics).Where("app_id = ?", app.ID)
	for queryKey, queryValue := range queries {
		startingQuery = addWhereQuery(startingQuery, queryKey, queryValue)
	}
	err := startingQuery.Select()
	return metrics, err
}

func addWhereQuery(query *orm.Query, whereQueryKey string, whereQueryValue string) *orm.Query {
	return query.Where(fmt.Sprintf("%s = ?", whereQueryKey), whereQueryValue)
}
