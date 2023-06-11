package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type DatabaseProps struct {
	Enabled  bool          `json:"enabled"`
	MariaDB  MariaDBProps  `json:"mariadb"`
	Postgres PostgresProps `json:"postgres"`
	Redis    RedisProps    `json:"redis"`
}

func NewDatabase(scope packager.Construct, props DatabaseProps) packager.Construct {
	if !props.Enabled {
		return nil
	}
	defer logModuleTiming("database")()

	chart := k8sapp.NewNamespaceChart(scope, "database")

	NewMariaDB(chart, props.MariaDB)
	NewPostgres(chart, props.Postgres)
	NewRedis(chart, props.Redis)

	return chart
}
