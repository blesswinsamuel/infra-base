package k8sbase

import (
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"github.com/muesli/reflow/dedent"
	corev1 "k8s.io/api/core/v1"
)

type KubeGitOpsProps struct {
	ImageInfo k8sapp.ImageInfo `json:"image"`
	GitRepo   struct {
		URL    string `json:"url"`
		Branch string `json:"branch"`
	} `json:"gitRepo"`
	GitSync struct {
		ImageInfo k8sapp.ImageInfo `json:"image"`
	} `json:"gitSync"`
	Webhook struct {
		ImageInfo k8sapp.ImageInfo `json:"image"`
	} `json:"webhook"`
	Kapp struct {
		AppName   string `json:"appName"`
		Directory string `json:"directory"`
		Namespace string `json:"namespace"`
	} `json:"kapp"`
}

func (props *KubeGitOpsProps) Render(scope kubegogen.Scope) {
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name: "kube-gitops",
		Containers: []k8sapp.ApplicationContainer{
			{
				Name:  "webhook",
				Image: props.Webhook.ImageInfo,
				Ports: []k8sapp.ContainerPort{{Name: "http", Port: 80}},
				Args: []string{
					"webhook",
					"-hooks=/config/hooks.yaml",
					"-verbose",
					"-port=80",
					"-hotreload",
				},
				Env: map[string]string{
					"GIT_REPO_PATH": "/repo/helm-chart-git/current",
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "git-repo", MountPath: "/repo", ReadOnly: true},
				},
			},
			{
				Name:  "git-sync",
				Image: props.GitSync.ImageInfo,
				Args: []string{
					"--repo=" + props.GitRepo.URL,
					"--branch=" + props.GitRepo.Branch,
					"--ssh=true",
					"--ssh-key-file=/ssh-key/git-private-key",
					"--ssh-known-hosts=false",
					"--ssh-known-hosts-file=/ssh-key/known_hosts",
					"--depth=1",
					"--max-sync-failures=200", // --max-failures=200
					"--wait=60",               // --period=60s
					"--dest=current",
					"--root=/repo/helm-chart-git",
					"--webhook-url=http://localhost/hooks/redeploy-webhook",
					"--webhook-timeout=120s",
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "git-repo", MountPath: "/repo"},
				},
			},
		},
		ExtraVolumes: []corev1.Volume{
			{Name: "git-repo", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
		},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{
			{
				Name: "kube-gitops-git-sync",
				RemoteRefs: map[string]string{
					"git-private-key": "GITHUB_DEPLOY_KEY",
					"known_hosts":     "GITHUB_KNOWN_HOSTS",
				},
				MountToContainers: []string{"git-sync"},
				MountName:         "git-keys",
				MountPath:         "/ssh-key",
				ReadOnly:          true,
			},
		},
		ConfigMaps: []k8sapp.ApplicationConfigMap{
			{
				Name: "kube-gitops-webhook-config",
				Data: map[string]string{
					"hooks.yaml": infrahelpers.ToYamlString([]map[string]any{
						{
							"id":                        "redeploy-webhook",
							"execute-command":           "/config/redeploy.sh",
							"command-working-directory": "/config",
						},
					}),
				},
				MountToContainers: []string{"webhook"},
				MountName:         "webhook-config",
				MountPath:         "/config",
				ReadOnly:          true,
			},
			{
				Name: "kube-gitops-scripts",
				Data: map[string]string{
					"redeploy.sh": strings.TrimSpace(dedent.String(`
						#!/bin/bash
						set -ex

						HELM_CHART_DIR=$GIT_REPO_PATH/{{ .Values.kubeChartPath }}

						{{- .Values.scripts.predeploy | nindent 4 }}

						kube -n {{ .Values.kubeReleaseNamespace }} diff upgrade {{ .Values.kubeReleaseName }} {{ $extraArgs }} $HELM_CHART_DIR --three-way-merge
						kube -n {{ .Values.kubeReleaseNamespace }} upgrade --install {{ .Values.kubeReleaseName }} {{ $extraArgs }} $HELM_CHART_DIR
					`)),
				},
				MountToContainers: []string{"webhook"},
				MountName:         "webhook-config",
				MountPath:         "/config",
				ReadOnly:          true,
			},
		},
	})
}

// gitRepo:
// url: "{{ .gitRepo.url }}"
// branch: "{{ .gitRepo.branch }}"
// kubeChartPath: "{{ .kubeChartPath }}"
// kubeReleaseName: "{{ .kubeReleaseName }}"
// kubeReleaseNamespace: "{{ .kubeReleaseNamespace }}"
// kubeExtraValuesFiles: {{ .kubeExtraValuesFiles | toJson }}
// clusterName: "test"
