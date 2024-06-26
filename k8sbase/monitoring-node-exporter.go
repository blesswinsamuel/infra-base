package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
)

type NodeExporterProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
	Disable       bool             `json:"disable"`
}

// https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus-node-exporter/values.yaml
func (props *NodeExporterProps) Render(scope kgen.Scope) {
	if props.Disable {
		return
	}
	// TODO: remove helm dependency
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "node-exporter",
		Values: map[string]interface{}{
			"fullnameOverride": "node-exporter",
			"service": map[string]interface{}{
				"annotations": map[string]string{
					"prometheus.io/scrape": "true",
					"prometheus.io/port":   "9100",
				},
			},
			"extraArgs": []string{
				"--collector.filesystem.mount-points-exclude=^/(dev|proc|run/credentials/.+|sys|var/lib/docker/.+|var/lib/containers/storage/.+|var/lib/kubelet/pods/.+|run/containerd/runc/k8s.io/.+)($|/)",
				"--collector.filesystem.fs-types-exclude=^(autofs|binfmt_misc|bpf|cgroup2?|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|iso9660|mqueue|nsfs|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|selinuxfs|squashfs|sysfs|tracefs|tmpfs)$",
			},
		},
	})
}
