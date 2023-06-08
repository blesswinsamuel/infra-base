package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type DatabaseProps struct {
	Namespace string        `yaml:"namespace"`
	MariaDB   MariaDBProps  `yaml:"mariadb"`
	Postgres  PostgresProps `yaml:"postgres"`
	Redis     RedisProps    `yaml:"redis"`
}

func NewDatabase(scope constructs.Construct, props DatabaseProps) constructs.Construct {
	defer logModuleTiming("database")()
	construct := constructs.NewConstruct(scope, jsii.String("database"))

	if props.Namespace != "" {
		k8sapp.SetNamespaceContext(construct, props.Namespace)
	} else {
		k8sapp.NewNamespaceChart(construct, "database")
	}

	NewMariaDB(construct, props.MariaDB)
	NewPostgres(construct, props.Postgres)
	NewRedis(construct, props.Redis)

	return construct
}
