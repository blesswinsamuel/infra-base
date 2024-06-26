package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
)

type GrafanaDashboardsProps struct {
	Dashboards infrahelpers.MergeableMap[string, k8sapp.GrafanaDashboard] `json:"dashboards"`
}

func (props *GrafanaDashboardsProps) Render(scope kgen.Scope) {
	k8sapp.NewGrafanaDashboards(scope, props.Dashboards)
}

// # [x] alertmanager-overview.json - "Alertmanager / Overview"
// # [x] apiserver.json - "Kubernetes / API server"
// # [ ] cluster-total.json - "Kubernetes / Networking / Cluster"
// # [ ] controller-manager.json - "Kubernetes / Controller Manager"
// # [x] grafana-overview.json - "Grafana Overview"
// # [ ] k8s-resources-cluster.json - "Kubernetes / Compute Resources / Cluster"
// # [ ] k8s-resources-namespace.json - "Kubernetes / Compute Resources / Namespace (Pods)"
// # [ ] k8s-resources-node.json - "Kubernetes / Compute Resources / Node (Pods)"
// # [ ] k8s-resources-pod.json - "Kubernetes / Compute Resources / Pod"
// # [ ] k8s-resources-workload.json - "Kubernetes / Compute Resources / Workload"
// # [ ] k8s-resources-workloads-namespace.json - "Kubernetes / Compute Resources / Namespace (Workloads)"
// # [ ] kubelet.json - "Kubernetes / Kubelet"
// # [ ] namespace-by-pod.json - "Kubernetes / Networking / Namespace (Pods)"
// # [ ] namespace-by-workload.json - "Kubernetes / Networking / Namespace (Workload)"
// # [ ] node-cluster-rsrc-use.json - "Node Exporter / USE Method / Cluster"
// # [ ] node-rsrc-use.json - "Node Exporter / USE Method / Node"
// # [x] nodes-darwin.json - "Node Exporter / MacOS"
// # [x] nodes.json - "Node Exporter / Nodes"
// # [x] persistentvolumesusage.json - "Kubernetes / Persistent Volumes"
// # [ ] pod-total.json - "Kubernetes / Networking / Pod"
// # [ ] proxy.json - "Kubernetes / Proxy"
// # [ ] scheduler.json - "Kubernetes / Scheduler"
// # [x] victoriametrics.json - "VictoriaMetrics"
// # [x] vmagent.json - "vmagent"
// # [x] vmalert.json - "vmalert"
// # [ ] workload-total.json - "Kubernetes / Networking / Workload"

// # {
// #     'source': 'https://raw.githubusercontent.com/prometheus-operator/kube-prometheus/main/manifests/grafana-dashboardDefinitions.yaml',
// #     'destination': '../dashboards/monitoring',
// #     'type': 'yaml',
// #     'dashboards': [
// #         'alertmanager-overview',
// #         'apiserver',
// #         'cluster-total',
// #         'controller-manager',
// #         'grafana-overview',
// #         'k8s-resources-cluster',
// #         'k8s-resources-namespace',
// #         'k8s-resources-node',
// #         'k8s-resources-pod',
// #         'k8s-resources-workload',
// #         'k8s-resources-workloads-namespace',
// #         'kubelet',
// #         'namespace-by-pod',
// #         'namespace-by-workload',
// #         'node-cluster-rsrc-use',
// #         'node-rsrc-use',
// #         'nodes-darwin',
// #         'nodes',
// #         'persistentvolumesusage',
// #         'pod-total',
// #         'prometheus-remote-write',
// #         'prometheus',
// #         'proxy',
// #         'scheduler',
// #         'workload-total',
// #     ],
// # },
