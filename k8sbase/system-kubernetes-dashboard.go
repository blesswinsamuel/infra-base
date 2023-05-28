package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/infraglobal"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type KubernetesDashboardProps struct {
	Enabled       bool              `yaml:"enabled"`
	HelmChartInfo helpers.ChartInfo `yaml:"helm"`
	Ingress       struct {
		Enabled   bool   `yaml:"enabled"`
		SubDomain string `yaml:"subDomain"`
	} `yaml:"ingress"`
}

// https://github.com/kubernetes/dashboard/tree/master/charts/helm-chart/kubernetes-dashboard
func NewKubernetesDashboard(scope constructs.Construct, props KubernetesDashboardProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: helpers.GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("kubernetes-dashboard"), &cprops)

	helpers.NewHelmCached(chart, jsii.String("helm"), &helpers.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("kubernetes-dashboard"),
		Namespace:   chart.Namespace(),
		Values: &map[string]any{
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
				"annotations": infraglobal.GetCertIssuerAnnotation(scope),
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

	// k8s.NewKubeServiceAccount(chart, jsii.String("service-account"), &k8s.KubeServiceAccountProps{
	// 	Metadata: &k8s.ObjectMeta{
	// 		Name:      jsii.String("kubernetes-dashboard"),
	// 		Namespace: chart.Namespace(),
	// 	},
	// })

	// k8s.NewKubeClusterRoleBinding(chart, jsii.String("cluster-role-binding"), &k8s.KubeClusterRoleBindingProps{
	// 	Metadata: &k8s.ObjectMeta{
	// 		Name: jsii.String("kubernetes-dashboard"),
	// 	},
	// 	RoleRef: &k8s.RoleRef{
	// 		ApiGroup: jsii.String("rbac.authorization.k8s.io"),
	// 		Kind:     jsii.String("ClusterRole"),
	// 		Name:     jsii.String("cluster-admin"),
	// 	},
	// 	Subjects: &[]*k8s.Subject{
	// 		{
	// 			Kind:      jsii.String("ServiceAccount"),
	// 			Name:      jsii.String("kubernetes-dashboard"),
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
