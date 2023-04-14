package resourcesbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DatabaseProps struct {
	MariaDB  MariaDBProps  `yaml:"mariadb"`
	Postgres PostgresProps `yaml:"postgres"`
	Redis    RedisProps    `yaml:"redis"`
}

func NewDatabase(scope constructs.Construct, props DatabaseProps) constructs.Construct {
	construct := constructs.NewConstruct(scope, jsii.String("database"))

	NewNamespace(construct, "database")

	NewMariaDB(construct, props.MariaDB)
	NewPostgres(construct, props.Postgres)
	NewRedis(construct, props.Redis)

	return construct
}
