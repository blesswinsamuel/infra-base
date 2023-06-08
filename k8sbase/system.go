package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type SystemProps struct {
	Reloader            ReloaderProps            `yaml:"reloader"`
	KubernetesDashboard KubernetesDashboardProps `yaml:"kubernetesDashboard"`
	Kopia               KopiaProps               `yaml:"kopia"`
	BackupJob           BackupJobProps           `yaml:"backupJob"`
}

func NewSystem(scope constructs.Construct, props SystemProps) constructs.Construct {
	defer logModuleTiming("system")()

	chart := k8sapp.NewNamespaceChart(scope, "system")

	NewReloader(chart, props.Reloader)
	NewKubernetesDashboard(chart, props.KubernetesDashboard)
	NewKopia(chart, props.Kopia)
	NewBackupJob(chart, props.BackupJob)

	return chart
}
