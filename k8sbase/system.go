package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type SystemProps struct {
	Reloader            ReloaderProps            `json:"reloader"`
	KubernetesDashboard KubernetesDashboardProps `json:"kubernetesDashboard"`
	Kopia               KopiaProps               `json:"kopia"`
	BackupJob           BackupJobProps           `json:"backupJob"`

	HelmOps map[string]interface{} `json:"helmOps"` // not implemented
}

func NewSystem(scope packager.Construct, props SystemProps) packager.Construct {
	defer logModuleTiming("system")()

	chart := k8sapp.NewNamespaceChart(scope, "system")

	NewReloader(chart, props.Reloader)
	NewKubernetesDashboard(chart, props.KubernetesDashboard)
	NewKopia(chart, props.Kopia)
	NewBackupJob(chart, props.BackupJob)

	return chart
}
