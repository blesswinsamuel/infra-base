package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type DatabaseProps struct {
	Enabled  bool          `yaml:"enabled"`
	MariaDB  MariaDBProps  `yaml:"mariadb"`
	Postgres PostgresProps `yaml:"postgres"`
	Redis    RedisProps    `yaml:"redis"`
}

func NewDatabase(scope constructs.Construct, props DatabaseProps) constructs.Construct {
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
