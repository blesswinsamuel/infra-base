package k8sapp

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/kgen"
	"golang.org/x/exp/slices"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

type ImageInfo struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

func (i ImageInfo) String() string {
	return fmt.Sprintf("%s:%s", i.Repository, i.Tag)
}

func (i *ImageInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"repository": i.Repository,
		"tag":        i.Tag,
	}
}

type ApplicationProps struct {
	Kind                         string
	Name                         string
	ServiceAccountName           string
	Replicas                     *int32
	CreateServiceAccount         bool
	Hostname                     string
	CreateHeadlessService        bool
	EnableServiceLinks           *bool
	AutomountServiceAccountToken bool
	AppAnnotations               map[string]string
	PodAnnotations               map[string]string
	PodSecurityContext           *corev1.PodSecurityContext
	ImagePullSecrets             string
	InitContainers               []ApplicationContainer
	Containers                   []ApplicationContainer
	ConfigMaps                   []ApplicationConfigMap
	ExternalSecrets              []ApplicationExternalSecret
	Secrets                      []ApplicationSecret
	PersistentVolumes            []ApplicationPersistentVolume // TODO: change to PersistentVolumeClaims
	ExtraVolumes                 []corev1.Volume
	HostNetwork                  bool
	DNSPolicy                    corev1.DNSPolicy
	DNSConfig                    *corev1.PodDNSConfig
	IngressMiddlewares           []NameNamespace
	IngressUseDefaultCert        *bool
	// IngressAnnotations              map[string]string

	DeploymentUpdateStrategy        v1.DeploymentStrategy
	StatefulSetUpdateStrategy       v1.StatefulSetUpdateStrategy
	StatefulSetServiceName          string
	StatefulSetVolumeClaimTemplates []ApplicationPersistentVolume
}

type ApplicationPersistentVolume struct {
	Name            string
	StorageClass    string
	VolumeName      string
	RequestsStorage string

	MountToContainers []string
	MountName         string
	MountPath         string
	SubPath           string
	ReadOnly          bool
}

type ApplicationConfigMap struct {
	Name string
	Data map[string]string

	MountToContainers []string
	MountName         string
	MountPath         string
	SubPath           string
	ReadOnly          bool
	DefaultMode       *int32
}

type ApplicationExternalSecret struct {
	Name       string
	RemoteRefs map[string]string
	Template   map[string]string

	MountToContainers []string
	MountName         string
	MountPath         string
	SubPath           string
	ReadOnly          bool

	VolumeItems []corev1.KeyToPath
}

type ApplicationSecret struct {
	Name string
	Data map[string]string

	MountToContainers []string
	MountName         string
	MountPath         string
	SubPath           string
	ReadOnly          bool

	VolumeItems []corev1.KeyToPath
}

type ApplicationContainer struct {
	Name              string
	Image             ImageInfo
	Command           []string
	Env               map[string]string
	EnvFromSecretRef  []string
	ExtraEnvs         []corev1.EnvVar
	Args              []string
	Ports             []ContainerPort
	ExtraVolumeMounts []corev1.VolumeMount
	SecurityContext   *corev1.SecurityContext
	LivenessProbe     *corev1.Probe
	ReadinessProbe    *corev1.Probe
	StartupProbe      *corev1.Probe
	Resources         corev1.ResourceRequirements
}

type ContainerPort struct {
	Name             string
	Port             int
	ServicePort      int
	Protocol         corev1.Protocol
	Ingress          *ApplicationIngress
	Ingresses        []ApplicationIngress
	PrometheusScrape *ApplicationPrometheusScrape
}

func (p ContainerPort) GetServicePort() int32 {
	if p.ServicePort != 0 {
		return int32(p.ServicePort)
	}
	return int32(p.Port)
}

type ApplicationIngress struct {
	Host string
	Path string // defaults to "/"
	TLS  *bool
}

type ApplicationPrometheusScrape struct {
	Path string // defaults to "/metrics"
}

func NewApplication(scope kgen.Scope, props *ApplicationProps) {
	if props.Kind == "" {
		props.Kind = "Deployment"
	}
	commonLabels := map[string]string{"app.kubernetes.io/name": props.Name}
	podAnnotations := infrahelpers.CopyMap(props.PodAnnotations)
	var volumes []corev1.Volume
	containerVolumeMountsMap := map[string][]corev1.VolumeMount{}
	allContainerNames := []string{}
	for _, container := range props.Containers {
		allContainerNames = append(allContainerNames, container.Name)
	}
	for _, container := range props.InitContainers {
		allContainerNames = append(allContainerNames, container.Name)
	}
	addVolumeMount := func(containerNames []string, mountName string, mountPath string, subPath string, readOnly bool) {
		if mountPath != "" {
			if containerNames == nil {
				containerNames = allContainerNames
			}
			for _, containerName := range containerNames {
				if _, ok := containerVolumeMountsMap[containerName]; !ok {
					containerVolumeMountsMap[containerName] = []corev1.VolumeMount{}
				}
				containerVolumeMountsMap[containerName] = append(containerVolumeMountsMap[containerName], corev1.VolumeMount{
					Name:      mountName,
					MountPath: mountPath,
					SubPath:   subPath,
					ReadOnly:  readOnly,
				})
			}
		}
	}
	watchTheseSecretsAndReload := []string{}
	addConfigMapHash := false
	configmapHash := sha256.New()
	for _, configmap := range props.ConfigMaps {
		NewConfigMap(scope, &ConfigmapProps{
			Name: configmap.Name,
			Data: configmap.Data,
		})
		if configmap.MountName == "" {
			continue
		}
		volumes = append(volumes, corev1.Volume{
			Name: configmap.MountName,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: configmap.Name}, DefaultMode: configmap.DefaultMode},
			},
		})
		addVolumeMount(configmap.MountToContainers, configmap.MountName, configmap.MountPath, configmap.SubPath, configmap.ReadOnly)
		for _, key := range infrahelpers.MapKeysSorted(configmap.Data) {
			configmapHash.Write([]byte(configmap.Data[key]))
			addConfigMapHash = true
		}
	}
	if addConfigMapHash {
		podAnnotations["configmap/checksum"] = fmt.Sprintf("%x", configmapHash.Sum(nil))
	}
	for _, secret := range props.Secrets {
		NewSecret(scope, &SecretProps{
			Name:       secret.Name,
			StringData: secret.Data,
		})
		if secret.MountName == "" {
			continue
		}
		volumes = append(volumes, corev1.Volume{
			Name:         secret.MountName,
			VolumeSource: corev1.VolumeSource{Secret: &corev1.SecretVolumeSource{SecretName: secret.Name, Items: secret.VolumeItems}},
		})
		addVolumeMount(secret.MountToContainers, secret.MountName, secret.MountPath, secret.SubPath, secret.ReadOnly)
		if secret.MountPath != "" {
			watchTheseSecretsAndReload = append(watchTheseSecretsAndReload, secret.Name)
		}
	}
	for _, externalSecret := range props.ExternalSecrets {
		if externalSecret.MountName != "" {
			volumes = append(volumes, corev1.Volume{
				Name:         externalSecret.MountName,
				VolumeSource: corev1.VolumeSource{Secret: &corev1.SecretVolumeSource{SecretName: externalSecret.Name, Items: externalSecret.VolumeItems}},
			})
			addVolumeMount(externalSecret.MountToContainers, externalSecret.MountName, externalSecret.MountPath, externalSecret.SubPath, externalSecret.ReadOnly)
			if externalSecret.MountPath != "" {
				watchTheseSecretsAndReload = append(watchTheseSecretsAndReload, externalSecret.Name)
			}
		}
		NewExternalSecret(scope, &ExternalSecretProps{
			Name:       externalSecret.Name,
			RemoteRefs: externalSecret.RemoteRefs,
			Template:   externalSecret.Template,
		})
	}
	volumes = append(volumes, props.ExtraVolumes...)
	for _, pv := range props.PersistentVolumes {
		if pv.MountName != "" {
			volumes = append(volumes, corev1.Volume{
				Name:         pv.MountName,
				VolumeSource: corev1.VolumeSource{PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{ClaimName: pv.Name}},
			})
			addVolumeMount(pv.MountToContainers, pv.MountName, pv.MountPath, pv.SubPath, pv.ReadOnly)
		}
		NewPersistentVolumeClaim(scope, &PersistentVolumeClaim{
			Name:            pv.Name,
			StorageClass:    pv.StorageClass,
			VolumeName:      pv.VolumeName,
			RequestsStorage: pv.RequestsStorage,
		})
	}
	statefulSetVolumeClaimTemplates := []corev1.PersistentVolumeClaim{}
	for _, pv := range props.StatefulSetVolumeClaimTemplates {
		if pv.MountName != "" {
			addVolumeMount(pv.MountToContainers, pv.MountName, pv.MountPath, pv.SubPath, pv.ReadOnly)
		}
		statefulSetVolumeClaimTemplates = append(statefulSetVolumeClaimTemplates, NewPersistentVolumeClaimProps(&PersistentVolumeClaim{
			Name:            pv.Name,
			StorageClass:    pv.StorageClass,
			VolumeName:      pv.VolumeName,
			RequestsStorage: pv.RequestsStorage,
		}))
	}
	containers := []corev1.Container{}
	initContainers := []corev1.Container{}
	servicePorts := []corev1.ServicePort{}
	ingressHosts := []IngressHost{}
	serviceAnnotations := map[string]string{}
	for _, container := range props.InitContainers {
		var containerVolumeMounts []corev1.VolumeMount
		containerVolumeMounts = append(containerVolumeMounts, containerVolumeMountsMap[container.Name]...)
		containerVolumeMounts = append(containerVolumeMounts, container.ExtraVolumeMounts...)
		env := []corev1.EnvVar{}
		env = append(env, container.ExtraEnvs...)
		for k, v := range container.Env {
			env = append(env, corev1.EnvVar{Name: k, Value: v})
		}
		envFrom := []corev1.EnvFromSource{}
		for _, v := range container.EnvFromSecretRef {
			envFrom = append(envFrom, corev1.EnvFromSource{
				SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: v}},
			})
		}
		initContainers = append(initContainers, corev1.Container{
			Name:    container.Name,
			Image:   container.Image.String(),
			Command: container.Command,
			// ImagePullPolicy: ("IfNotPresent"),
			Env:                      env,
			EnvFrom:                  envFrom,
			Args:                     container.Args,
			VolumeMounts:             containerVolumeMounts,
			LivenessProbe:            container.LivenessProbe,
			ReadinessProbe:           container.ReadinessProbe,
			StartupProbe:             container.StartupProbe,
			Resources:                container.Resources,
			TerminationMessagePolicy: "FallbackToLogsOnError",
			SecurityContext:          container.SecurityContext,
		})
	}
	for _, container := range props.Containers {
		var containerVolumeMounts []corev1.VolumeMount
		containerVolumeMounts = append(containerVolumeMounts, containerVolumeMountsMap[container.Name]...)
		containerVolumeMounts = append(containerVolumeMounts, container.ExtraVolumeMounts...)

		env := []corev1.EnvVar{}
		env = append(env, container.ExtraEnvs...)
		for k, v := range container.Env {
			env = append(env, corev1.EnvVar{Name: k, Value: v})
		}
		slices.SortFunc(env, func(a corev1.EnvVar, b corev1.EnvVar) int {
			return strings.Compare(a.Name, b.Name)
		})
		envFrom := []corev1.EnvFromSource{}
		for _, v := range container.EnvFromSecretRef {
			envFrom = append(envFrom, corev1.EnvFromSource{
				SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: v}},
			})
			watchTheseSecretsAndReload = append(watchTheseSecretsAndReload, v)
		}
		var ports []corev1.ContainerPort
		for _, port := range container.Ports {
			ports = append(ports, corev1.ContainerPort{
				Name:          port.Name,
				ContainerPort: int32(port.Port),
				Protocol:      port.Protocol,
			})
		}

		for _, port := range container.Ports {
			servicePorts = append(servicePorts, corev1.ServicePort{
				Name:       port.Name,
				Port:       port.GetServicePort(),
				TargetPort: intstr.FromString(port.Name),
				Protocol:   port.Protocol,
			})
			if port.Ingress != nil {
				port.Ingresses = append(port.Ingresses, *port.Ingress)
			}
			for _, ingress := range port.Ingresses {
				if ingress.TLS == nil {
					ingress.TLS = infrahelpers.Ptr(true)
				}
				ingressHosts = append(ingressHosts, IngressHost{
					Host:  ingress.Host,
					Paths: []IngressHostPath{{Path: ingress.Path, ServiceName: props.Name, ServicePortName: port.Name}},
					Tls:   *ingress.TLS,
				})
			}
			if prometheusScrape := port.PrometheusScrape; prometheusScrape != nil {
				serviceAnnotations["prometheus.io/scrape"] = "true"
				serviceAnnotations["prometheus.io/port"] = fmt.Sprint(port.GetServicePort())
				if prometheusScrape.Path != "" {
					serviceAnnotations["prometheus.io/path"] = prometheusScrape.Path
				}
			}
		}

		containers = append(containers, corev1.Container{
			Name:    container.Name,
			Image:   container.Image.String(),
			Command: container.Command,
			// ImagePullPolicy: ("IfNotPresent"),
			Env:                      env,
			EnvFrom:                  envFrom,
			Args:                     container.Args,
			VolumeMounts:             containerVolumeMounts,
			Ports:                    ports,
			LivenessProbe:            container.LivenessProbe,
			ReadinessProbe:           container.ReadinessProbe,
			StartupProbe:             container.StartupProbe,
			TerminationMessagePolicy: "FallbackToLogsOnError",
			SecurityContext:          container.SecurityContext,
			Resources:                container.Resources,
		})
	}
	for _, vol := range props.ExtraVolumes {
		if vol.Secret != nil {
			watchTheseSecretsAndReload = append(watchTheseSecretsAndReload, vol.Secret.SecretName)
		}
	}
	appAnnotations := infrahelpers.CopyMap(props.AppAnnotations)
	if len(watchTheseSecretsAndReload) > 0 {
		appAnnotations["secret.reloader.stakater.com/reload"] = strings.Join(watchTheseSecretsAndReload, ",")
	}

	podTemplate := corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      commonLabels,
			Annotations: podAnnotations,
		},
		Spec: corev1.PodSpec{
			ServiceAccountName:           props.ServiceAccountName,
			AutomountServiceAccountToken: infrahelpers.PtrIfNonEmpty(props.AutomountServiceAccountToken),
			Hostname:                     props.Hostname,
			EnableServiceLinks:           props.EnableServiceLinks,
			SecurityContext:              props.PodSecurityContext,
			ImagePullSecrets: infrahelpers.If(props.ImagePullSecrets != "", []corev1.LocalObjectReference{
				{Name: props.ImagePullSecrets},
			}, nil),
			Containers:     containers,
			InitContainers: initContainers,
			Volumes:        volumes,
			HostNetwork:    props.HostNetwork,
			DNSPolicy:      props.DNSPolicy,
			DNSConfig:      props.DNSConfig,
		},
	}
	if props.Replicas == nil {
		props.Replicas = infrahelpers.Ptr(int32(1))
	}
	switch props.Kind {
	case "Deployment":
		scope.AddApiObject(&v1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:        props.Name,
				Annotations: appAnnotations,
			},
			Spec: v1.DeploymentSpec{
				Replicas: props.Replicas,
				Strategy: props.DeploymentUpdateStrategy,
				Selector: &metav1.LabelSelector{
					MatchLabels: commonLabels,
				},
				Template: podTemplate,
			},
		})
	case "StatefulSet":
		scope.AddApiObject(&v1.StatefulSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:        props.Name,
				Annotations: appAnnotations,
			},
			Spec: v1.StatefulSetSpec{
				Replicas:       props.Replicas,
				UpdateStrategy: props.StatefulSetUpdateStrategy,
				ServiceName:    infrahelpers.UseOrDefault(props.StatefulSetServiceName, props.Name),
				Selector: &metav1.LabelSelector{
					MatchLabels: commonLabels,
				},
				Template:             podTemplate,
				VolumeClaimTemplates: statefulSetVolumeClaimTemplates,
			},
		})
	case "DaemonSet":
		scope.AddApiObject(&v1.DaemonSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:        props.Name,
				Annotations: appAnnotations,
			},
			Spec: v1.DaemonSetSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: commonLabels,
				},
				Template: podTemplate,
			},
		})
	}
	if len(servicePorts) > 0 {
		scope.AddApiObject(&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:        props.Name,
				Annotations: serviceAnnotations,
			},
			Spec: corev1.ServiceSpec{
				Selector: commonLabels,
				Ports:    servicePorts,
			},
		})
		if props.CreateHeadlessService {
			scope.AddApiObject(&corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: props.Name + "-headless",
				},
				Spec: corev1.ServiceSpec{
					ClusterIP: "None",
					Selector:  commonLabels,
					Ports:     servicePorts,
				},
			})
		}
	}
	if len(ingressHosts) > 0 {
		if props.IngressUseDefaultCert == nil {
			props.IngressUseDefaultCert = ptr.To(true)
		}
		NewIngress(scope, &IngressProps{
			Name:                   props.Name,
			Hosts:                  ingressHosts,
			TraefikMiddlewareNames: props.IngressMiddlewares,
			UseDefaultCert:         *props.IngressUseDefaultCert,
			// Annotations:            props.IngressAnnotations,
		})
	}
	if props.CreateServiceAccount {
		scope.AddApiObject(&corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name: props.ServiceAccountName,
			},
		})
	}
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
