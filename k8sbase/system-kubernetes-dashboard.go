package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type KubernetesDashboardProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
	Ingress       struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
}

// https://github.com/kubernetes/dashboard/tree/master/charts/helm-chart/kubernetes-dashboard
func (props *KubernetesDashboardProps) Chart(scope packager.Construct) packager.Construct {
	chart := k8sapp.NewHelmChart(scope, "kubernetes-dashboard", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "kubernetes-dashboard",
		Values: map[string]any{
			"protocolHttp": true,
			"extraArgs": []string{
				"--enable-skip-login",
				"--enable-insecure-login",
			},
			"service": map[string]any{
				"externalPort": 9090,
			},
			"rbac": map[string]any{
				"create":              true,
				"clusterReadOnlyRole": true,
			},
			"serviceAccount": map[string]any{
				"create": true,
				"name":   "kubernetes-dashboard",
			},
			"metricsScraper": map[string]any{
				"enabled": true,
			},
			"ingress": map[string]any{
				"enabled":     props.Ingress.Enabled,
				"annotations": GetCertIssuerAnnotation(scope),
				"hosts": []string{
					props.Ingress.SubDomain + "." + GetDomain(scope),
				},
				"tls": []map[string]any{
					{
						"hosts": []string{
							props.Ingress.SubDomain + "." + GetDomain(scope),
						},
						"secretName": "kubernetes-dashboard-tls",
					},
				},
			},
		},
	})

	// k8s.NewKubeServiceAccount(chart, ("service-account"), &k8s.KubeServiceAccountProps{
	// 	Metadata: &k8s.ObjectMeta{
	// 		Name:      ("kubernetes-dashboard"),
	// 		Namespace: chart.Namespace(),
	// 	},
	// })

	// k8s.NewKubeClusterRoleBinding(chart, ("cluster-role-binding"), &k8s.KubeClusterRoleBindingProps{
	// 	Metadata: &k8s.ObjectMeta{
	// 		Name: ("kubernetes-dashboard"),
	// 	},
	// 	RoleRef: &k8s.RoleRef{
	// 		ApiGroup: ("rbac.authorization.k8s.io"),
	// 		Kind:     ("ClusterRole"),
	// 		Name:     ("cluster-admin"),
	// 	},
	// 	Subjects: &[]*k8s.Subject{
	// 		{
	// 			Kind:      ("ServiceAccount"),
	// 			Name:      ("kubernetes-dashboard"),
	// 			Namespace: chart.Namespace(),
	// 		},
	// 	},
	// })

	return chart
}

// # ---
// # apiVersion: v1
// # kind: Secret
// # metadata:
// #   name: kubernetes-dashboard
// #   namespace: '{{ tpl $.Values.system.namespace $ }}'
// #   annotations:
// #     kubernetes.io/service-account.name: kubernetes-dashboard
// # type: kubernetes.io/service-account-token
