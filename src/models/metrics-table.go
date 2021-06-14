package models

import "github.com/go-pg/pg/v10"

type MetricsTable struct {
	tableName       struct{}          `pg:"metrics"`
	ID              int               `pg:"id,pk" json:"id"`
	AppVersion      string            `pg:"app_version" json:"app_version"`
	AppBuildVersion string            `pg:"app_build_version" json:"app_build_version"`
	Payload         CollectionMetrics `pg:"payload" json:"payload"`
	AppID           int               `pg:"app_id" json:"app_id"`
}

func (metric *MetricsTable) Save(pgDB *pg.DB) error {
	_, err := pgDB.Model(metric).Insert()
	return err
}
