package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type SystemProps struct {
	Reloader            ReloaderProps            `json:"reloader"`
	KubernetesDashboard KubernetesDashboardProps `json:"kubernetesDashboard"`
	KubeGitOps          KubeGitOpsProps          `json:"kubeGitOps"`
	Kopia               KopiaProps               `json:"kopia"`
	BackupJob           BackupJobProps           `json:"backupJob"`
}

func NewSystem(scope packager.Construct, props SystemProps) packager.Construct {
	defer logModuleTiming("system")()

	chart := k8sapp.NewNamespaceChart(scope, "system")

	NewReloader(chart, props.Reloader)
	NewKubernetesDashboard(chart, props.KubernetesDashboard)
	NewKubeGitOps(chart, props.KubeGitOps)
	NewKopia(chart, props.Kopia)
	NewBackupJob(chart, props.BackupJob)

	return chart
}
