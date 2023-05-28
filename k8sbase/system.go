package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
)

type SystemProps struct {
	Reloader            ReloaderProps            `yaml:"reloader"`
	KubernetesDashboard KubernetesDashboardProps `yaml:"kubernetesDashboard"`
	Kopia               KopiaProps               `yaml:"kopia"`
	BackupJob           BackupJobProps           `yaml:"backupJob"`
}

func NewSystem(scope constructs.Construct, props SystemProps) constructs.Construct {
	construct := constructs.NewConstruct(scope, jsii.String("system"))

	helpers.NewNamespace(construct, "system")

	NewReloader(construct, props.Reloader)
	NewKubernetesDashboard(construct, props.KubernetesDashboard)
	NewKopia(construct, props.Kopia)
	NewBackupJob(construct, props.BackupJob)

	return construct
}
