package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DatabaseProps struct {
	Namespace string        `yaml:"namespace"`
	MariaDB   MariaDBProps  `yaml:"mariadb"`
	Postgres  PostgresProps `yaml:"postgres"`
	Redis     RedisProps    `yaml:"redis"`
}

func NewDatabase(scope constructs.Construct, props DatabaseProps) constructs.Construct {
	construct := constructs.NewConstruct(scope, jsii.String("database"))

	if props.Namespace != "" {
		UseNamespace(construct, props.Namespace)
	} else {
		NewNamespace(construct, "database")
	}

	NewMariaDB(construct, props.MariaDB)
	NewPostgres(construct, props.Postgres)
	NewRedis(construct, props.Redis)

	return construct
}
