package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type SystemProps struct {
	Reloader            ReloaderProps            `json:"reloader"`
	KubernetesDashboard KubernetesDashboardProps `json:"kubernetesDashboard"`
	Kopia               KopiaProps               `json:"kopia"`
	BackupJob           BackupJobProps           `json:"backupJob"`

	HelmOps map[string]interface{} `json:"helmOps"` // not implemented
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
