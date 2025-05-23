package k8sapp

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/kgen"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

type ApplicationKind string

const (
	ApplicationKindDeployment  ApplicationKind = "Deployment"
	ApplicationKindStatefulSet                 = "StatefulSet"
	ApplicationKindDaemonSet                   = "DaemonSet"
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
	// Name is the name of the application. This is optional. If not provided, the name will be derived from the scope.
	Name string
	// Kind is the kind of the application. This is optional. If not provided, it will default to Deployment.
	// Valid values are Deployment, StatefulSet, and DaemonSet.
	Kind ApplicationKind

	ServiceAccountName            string
	Replicas                      *int32
	CreateServiceAccount          bool
	Hostname                      string
	HeadlessServiceNames          []string
	EnableServiceLinks            *bool
	AutomountServiceAccountToken  *bool
	AppAnnotations                map[string]string
	PodAnnotations                map[string]string
	PodSecurityContext            *corev1.PodSecurityContext
	ImagePullSecrets              string
	InitContainers                []ApplicationContainer
	Containers                    []ApplicationContainer
	ConfigMaps                    []ApplicationConfigMap
	ExternalSecrets               []ApplicationExternalSecret
	Secrets                       []ApplicationSecret
	PersistentVolumes             []ApplicationPersistentVolume // TODO: change to PersistentVolumeClaims
	ExtraVolumes                  []corev1.Volume
	HostNetwork                   bool
	HostPID                       bool
	NodeSelector                  map[string]string
	Tolerations                   []corev1.Toleration
	DNSPolicy                     corev1.DNSPolicy
	DNSConfig                     *corev1.PodDNSConfig
	IngressMiddlewares            []NameNamespace
	IngressUseDefaultCert         *bool
	Affinity                      *corev1.Affinity
	TerminationGracePeriodSeconds *int64
	Homepage                      *ApplicationHomepage
	NetworkPolicy                 *ApplicationNetworkPolicy
	Security                      *ApplicationSecurity
	// IngressAnnotations              map[string]string

	DeploymentUpdateStrategy        v1.DeploymentStrategy
	StatefulSetUpdateStrategy       v1.StatefulSetUpdateStrategy
	DaemonSetUpdateStrategy         v1.DaemonSetUpdateStrategy
	StatefulSetServiceName          string
	StatefulSetVolumeClaimTemplates []ApplicationPersistentVolume
}

type ApplicationSecurity struct {
	User                     int64
	Group                    int64
	FSGroup                  int64
	WriteableRootFS          bool
	AllowPrivilegeEscalation bool
	RunAsRoot                bool
	Capabilities             []corev1.Capability
}

type ApplicationNetworkPolicy struct {
	Ingress NetworkPolicyIngress
	Egress  NetworkPolicyEgress
}

type ApplicationHomepage struct {
	Name        string
	Description string
	Group       string
	Icon        string
	Widget      map[string]string
	SiteMonitor string
	Href        string
	PodSelector string
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
	AccessModes       []corev1.PersistentVolumeAccessMode
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
	Name            string
	Image           ImageInfo
	ImagePullPolicy corev1.PullPolicy

	Env              map[string]string
	EnvFromSecretRef []string
	ExtraEnvs        []corev1.EnvVar

	Command []string
	Args    []string

	SecurityContext *corev1.SecurityContext
	LivenessProbe   *corev1.Probe
	ReadinessProbe  *corev1.Probe
	StartupProbe    *corev1.Probe
	Resources       corev1.ResourceRequirements

	// ExtraVolumeMounts are additional volume mounts to be added to the container.
	//
	// Deprecated: use VolumeMounts instead
	ExtraVolumeMounts []corev1.VolumeMount
	// VolumeMounts are the volume mounts to be added to the container.
	VolumeMounts []corev1.VolumeMount

	// Deprecated
	Ports []ContainerPort
}

type ContainerPort struct {
	Name                 string
	Port                 int
	ServicePort          int
	NodePort             int32
	ServiceName          string
	DisableService       bool
	DisableContainerPort bool
	Protocol             corev1.Protocol
	Ingress              *ApplicationIngress
	Ingresses            []ApplicationIngress
	PrometheusScrape     *ApplicationPrometheusScrape
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
	if props.Name == "" {
		props.Name = scope.ID()
	}
	if props.Kind == "" {
		props.Kind = ApplicationKindDeployment
	}
	commonLabels := map[string]string{"app.kubernetes.io/name": props.Name}
	podAnnotations := maps.Clone(props.PodAnnotations)
	if podAnnotations == nil {
		podAnnotations = map[string]string{}
	}
	var volumes []corev1.Volume
	containerVolumeMountsMap := map[string][]corev1.VolumeMount{}
	allContainerNames := []string{}
	networkPolicy := NetworkPolicy{Name: props.Name}
	if props.NetworkPolicy != nil {
		networkPolicy.Ingress = props.NetworkPolicy.Ingress
		networkPolicy.Egress = props.NetworkPolicy.Egress
	}
	if networkPolicy.Ingress.AllowFromAppRefs == nil {
		networkPolicy.Ingress.AllowFromAppRefs = map[string][]intstr.IntOrString{}
	}
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
	if props.Security != nil {
		if props.PodSecurityContext == nil {
			props.PodSecurityContext = &corev1.PodSecurityContext{
				RunAsNonRoot:   ptr.To(true),
				SeccompProfile: &corev1.SeccompProfile{Type: corev1.SeccompProfileTypeRuntimeDefault},
			}
			if props.Security.RunAsRoot {
				props.PodSecurityContext.RunAsNonRoot = ptr.To(false)
			}
		}
		for _, container := range infrahelpers.MergeLists(infrahelpers.ListToPtrs(props.Containers), infrahelpers.ListToPtrs(props.InitContainers)) {
			if container.SecurityContext == nil {
				container.SecurityContext = &corev1.SecurityContext{
					AllowPrivilegeEscalation: ptr.To(false),
					ReadOnlyRootFilesystem:   ptr.To(true),
					Capabilities:             &corev1.Capabilities{Drop: []corev1.Capability{"ALL"}},
				}
			}
			if props.Security.WriteableRootFS {
				container.SecurityContext.ReadOnlyRootFilesystem = ptr.To(false)
			}
			if props.Security.AllowPrivilegeEscalation {
				container.SecurityContext.AllowPrivilegeEscalation = ptr.To(true)
			}
			if len(props.Security.Capabilities) > 0 {
				container.SecurityContext.Capabilities.Add = append(container.SecurityContext.Capabilities.Add, props.Security.Capabilities...)
			}
		}
		if props.Security.User != 0 {
			props.PodSecurityContext.RunAsUser = ptr.To(props.Security.User)
		}
		if props.Security.Group != 0 {
			props.PodSecurityContext.RunAsGroup = ptr.To(props.Security.Group)
		}
		if props.Security.FSGroup != 0 {
			props.PodSecurityContext.FSGroup = ptr.To(props.Security.FSGroup)
		}
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
			AccessModes:     pv.AccessModes,
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
			AccessModes:     pv.AccessModes,
		}))
	}
	containers := []corev1.Container{}
	initContainers := []corev1.Container{}
	servicePorts := map[string][]corev1.ServicePort{}
	ingressHosts := []IngressHost{}
	serviceAnnotations := map[string]map[string]string{}
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
			Name:                     container.Name,
			Image:                    container.Image.String(),
			Command:                  container.Command,
			ImagePullPolicy:          container.ImagePullPolicy,
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
			if port.DisableContainerPort {
				continue
			}
			ports = append(ports, corev1.ContainerPort{
				Name:          port.Name,
				ContainerPort: int32(port.Port),
				Protocol:      port.Protocol,
			})
		}

		for _, port := range container.Ports {
			if port.DisableService {
				continue
			}
			serviceName := props.Name
			if port.ServiceName != "" {
				serviceName = port.ServiceName
			}
			servicePorts[serviceName] = append(servicePorts[serviceName], corev1.ServicePort{
				Name:       port.Name,
				Port:       port.GetServicePort(),
				TargetPort: intstr.FromString(port.Name),
				Protocol:   port.Protocol,
				NodePort:   port.NodePort,
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
			if len(port.Ingresses) > 0 {
				networkPolicy.Ingress.AllowFromAppRefs["traefik"] = append(networkPolicy.Ingress.AllowFromAppRefs["traefik"], intstr.FromString(port.Name))
			}
			if prometheusScrape := port.PrometheusScrape; prometheusScrape != nil {
				if serviceAnnotations[serviceName] == nil {
					serviceAnnotations[serviceName] = map[string]string{}
				}
				serviceAnnotations[serviceName]["prometheus.io/scrape"] = "true"
				serviceAnnotations[serviceName]["prometheus.io/port"] = fmt.Sprint(port.GetServicePort())
				if prometheusScrape.Path != "" {
					serviceAnnotations[serviceName]["prometheus.io/path"] = prometheusScrape.Path
				}
				if len(networkPolicy.Ingress.AllowFromAppRefs["vmagent"]) == 0 {
					networkPolicy.Ingress.AllowFromAppRefs["vmagent"] = []intstr.IntOrString{intstr.FromInt32(port.GetServicePort())}
				}
			}
		}

		containers = append(containers, corev1.Container{
			Name:                     container.Name,
			Image:                    container.Image.String(),
			Command:                  container.Command,
			ImagePullPolicy:          container.ImagePullPolicy,
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
	appAnnotations := maps.Clone(props.AppAnnotations)
	if appAnnotations == nil {
		appAnnotations = map[string]string{}
	}
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
			AutomountServiceAccountToken: props.AutomountServiceAccountToken,
			Hostname:                     props.Hostname,
			EnableServiceLinks:           props.EnableServiceLinks,
			SecurityContext:              props.PodSecurityContext,
			ImagePullSecrets: infrahelpers.If(props.ImagePullSecrets != "", []corev1.LocalObjectReference{
				{Name: props.ImagePullSecrets},
			}, nil),
			Containers:                    containers,
			InitContainers:                initContainers,
			Volumes:                       volumes,
			HostNetwork:                   props.HostNetwork,
			HostPID:                       props.HostPID,
			NodeSelector:                  props.NodeSelector,
			Tolerations:                   props.Tolerations,
			DNSPolicy:                     props.DNSPolicy,
			DNSConfig:                     props.DNSConfig,
			Affinity:                      props.Affinity,
			TerminationGracePeriodSeconds: props.TerminationGracePeriodSeconds,
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
				UpdateStrategy: props.DaemonSetUpdateStrategy,
				Selector: &metav1.LabelSelector{
					MatchLabels: commonLabels,
				},
				Template: podTemplate,
			},
		})
	}
	if len(servicePorts) > 0 {
		for _, serviceName := range infrahelpers.MapKeysSorted(servicePorts) {
			clusterIP := ""
			if slices.Contains(props.HeadlessServiceNames, serviceName) {
				clusterIP = "None"
			}
			var serviceType corev1.ServiceType
			if strings.HasSuffix(serviceName, "-np") {
				serviceType = corev1.ServiceTypeNodePort
			}
			scope.AddApiObject(&corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:        serviceName,
					Annotations: serviceAnnotations[serviceName],
				},
				Spec: corev1.ServiceSpec{
					Type:      serviceType,
					ClusterIP: clusterIP,
					Selector:  commonLabels,
					Ports:     servicePorts[serviceName],
				},
			})
		}
	}
	if len(ingressHosts) > 0 {
		if props.IngressUseDefaultCert == nil {
			props.IngressUseDefaultCert = ptr.To(true)
		}
		var ingressAnnotations map[string]string
		if props.Homepage != nil {
			ingressAnnotations = GetHomepageAnnotations(props.Homepage)
			if len(networkPolicy.Ingress.AllowFromAppRefs["homepage"]) == 0 && (GetGlobals(scope).AppRefs["homepage"] != NameNamespacePort{}) {
				ports := []intstr.IntOrString{}
				for _, port := range ingressHosts {
					for _, path := range port.Paths {
						ports = append(ports, intstr.FromString(path.ServicePortName))
					}
				}
				networkPolicy.Ingress.AllowFromAppRefs["homepage"] = ports
			}
		}
		NewIngress(scope, &IngressProps{
			Name:                   props.Name,
			Hosts:                  ingressHosts,
			TraefikMiddlewareNames: props.IngressMiddlewares,
			UseDefaultCert:         *props.IngressUseDefaultCert,
			Annotations:            ingressAnnotations,
		})
	}
	if props.CreateServiceAccount {
		scope.AddApiObject(&corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name: props.ServiceAccountName,
			},
		})
	}
	if props.NetworkPolicy != nil {
		NewNetworkPolicy(scope, &networkPolicy)
	}
}

func GetHomepageAnnotations(props *ApplicationHomepage) map[string]string {
	if props == nil {
		return nil
	}
	annotations := map[string]string{
		"gethomepage.dev/enabled":     "true",
		"gethomepage.dev/description": props.Description,
		"gethomepage.dev/group":       props.Group,
		"gethomepage.dev/icon":        props.Icon,
		"gethomepage.dev/name":        props.Name,
		"gethomepage.dev/siteMonitor": props.SiteMonitor,
	}
	if props.Href != "" {
		annotations["gethomepage.dev/href"] = props.Href
	}
	if props.PodSelector != "" {
		annotations["gethomepage.dev/pod-selector"] = props.PodSelector
	}
	for key, value := range props.Widget {
		annotations[fmt.Sprintf("gethomepage.dev/widget.%s", key)] = value
	}
	return annotations
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
