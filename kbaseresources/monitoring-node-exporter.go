package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

func init() {
	RegisterModule("node-exporter", &NodeExporterProps{})
}

type NodeExporterProps struct {
	ImageInfo k8sapp.ImageInfo `json:"image"`
	Disable   bool             `json:"disable"`
}

// https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus-node-exporter/values.yaml
func (props *NodeExporterProps) Render(scope kgen.Scope) {
	if props.Disable {
		return
	}
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name:                 "node-exporter",
		ServiceAccountName:   "node-exporter",
		Kind:                 "DaemonSet",
		CreateServiceAccount: true,
		PodAnnotations: map[string]string{
			"cluster-autoscaler.kubernetes.io/safe-to-evict": "true",
		},
		AutomountServiceAccountToken: ptr.To(false),
		Containers: []k8sapp.ApplicationContainer{
			{
				Name: "node-exporter",
				Args: []string{
					"--path.procfs=/host/proc",
					"--path.sysfs=/host/sys",
					"--path.rootfs=/host/root",
					"--path.udev.data=/host/root/run/udev/data",
					"--web.listen-address=[$(HOST_IP)]:9100",
					"--collector.filesystem.mount-points-exclude=^/(dev|proc|run/credentials/.+|sys|var/lib/docker/.+|var/lib/containers/storage/.+|var/lib/kubelet/pods/.+|run/containerd/runc/k8s.io/.+)($|/)",
					"--collector.filesystem.fs-types-exclude=^(autofs|binfmt_misc|bpf|cgroup2?|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|iso9660|mqueue|nsfs|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|selinuxfs|squashfs|sysfs|tracefs|tmpfs)$",
				},
				Env:            map[string]string{"HOST_IP": "0.0.0.0"},
				Image:          props.ImageInfo,
				Ports:          []k8sapp.ContainerPort{{Name: "metrics", Port: 9100, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}}},
				LivenessProbe:  &corev1.Probe{FailureThreshold: 3, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/", Port: intstr.FromString("metrics")}}, PeriodSeconds: 10, SuccessThreshold: 1, TimeoutSeconds: 1},
				ReadinessProbe: &corev1.Probe{FailureThreshold: 3, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/", Port: intstr.FromString("metrics")}}, PeriodSeconds: 10, SuccessThreshold: 1, TimeoutSeconds: 1},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{MountPath: "/host/proc", Name: "proc", ReadOnly: true},
					{MountPath: "/host/sys", Name: "sys", ReadOnly: true},
					{MountPath: "/host/root", MountPropagation: ptr.To(corev1.MountPropagationHostToContainer), Name: "root", ReadOnly: true},
				},
			},
		},
		HostNetwork: true,
		HostPID:     true,
		NodeSelector: map[string]string{
			"kubernetes.io/os": "linux",
		},
		Tolerations: []corev1.Toleration{
			{Effect: corev1.TaintEffectNoSchedule, Operator: corev1.TolerationOpExists},
		},
		ExtraVolumes: []corev1.Volume{
			{Name: "proc", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/proc"}}},
			{Name: "sys", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/sys"}}},
			{Name: "root", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/"}}},
		},
		DaemonSetUpdateStrategy: appsv1.DaemonSetUpdateStrategy{Type: appsv1.RollingUpdateDaemonSetStrategyType, RollingUpdate: &appsv1.RollingUpdateDaemonSet{MaxUnavailable: ptr.To(intstr.FromInt(1))}},
		Security:                &k8sapp.ApplicationSecurity{User: 65534, Group: 65534, FSGroup: 65534},
		NetworkPolicy:           &k8sapp.ApplicationNetworkPolicy{},
	})
}
