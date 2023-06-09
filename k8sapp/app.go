package k8sapp

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"golang.org/x/exp/slices"
)

type ImageInfo struct {
	Repository string `yaml:"repository"`
	Tag        string `yaml:"tag"`
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
	Kind                     string
	Name                     string
	ServiceAccountName       string
	Hostname                 string
	EnableServiceLinks       bool
	DeploymentUpdateStrategy *k8s.DeploymentStrategy
	AppAnnotations           map[string]string
	PodAnnotations           map[string]string
	PodSecurityContext       *k8s.PodSecurityContext
	ImagePullSecrets         string
	Containers               []ApplicationContainer
	ConfigMaps               []ApplicationConfigMap
	ExternalSecrets          []ApplicationExternalSecret
	Secrets                  []ApplicationSecret
	PersistentVolumes        []ApplicationPersistentVolume
	ExtraVolumes             []*k8s.Volume
	HostNetwork              bool
	DnsPolicy                string
	IngressMiddlewares       []NameNamespace
	// IngressAnnotations       map[string]string
}

type ApplicationPersistentVolume struct {
	Name            string
	StorageClass    string
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
}

type ApplicationSecret struct {
	Name string
	Data map[string]string

	MountToContainers []string
	MountName         string
	MountPath         string
	SubPath           string
	ReadOnly          bool
}

type ApplicationContainer struct {
	Name              string
	Image             ImageInfo
	Command           []string
	Env               map[string]string
	EnvFromSecretRef  []string
	Args              []string
	Ports             []ContainerPort
	ExtraVolumeMounts []*k8s.VolumeMount
	SecurityContext   *k8s.SecurityContext
	LivenessProbe     *k8s.Probe
	ReadinessProbe    *k8s.Probe
	StartupProbe      *k8s.Probe
}

type ContainerPort struct {
	Name             string
	Port             int
	Ingress          *ApplicationIngress
	PrometheusScrape *ApplicationPrometheusScrape
}

type ApplicationIngress struct {
	Host string
	Path string // defaults to "/"
}

type ApplicationPrometheusScrape struct {
	Path string // defaults to "/metrics"
}

func NewApplicationChart(scope constructs.Construct, id string, props *ApplicationProps) cdk8s.Chart {
	chart := cdk8s.NewChart(scope, jsii.String(id), &cdk8s.ChartProps{
		Namespace: GetNamespaceContextPtr(scope),
	})
	NewApplication(chart, jsii.String("application"), props)
	return chart
}

func NewApplication(scope constructs.Construct, id *string, props *ApplicationProps) constructs.Construct {
	scope = constructs.NewConstruct(scope, id)
	if props.Kind == "" {
		props.Kind = "Deployment"
	}
	commonLabels := map[string]string{"app.kubernetes.io/name": props.Name}
	podAnnotations := infrahelpers.CopyMap(props.PodAnnotations)
	var volumes []*k8s.Volume
	containerVolumeMountsMap := map[string][]*k8s.VolumeMount{}
	allContainerNames := []string{}
	for _, container := range props.Containers {
		allContainerNames = append(allContainerNames, container.Name)
	}
	addVolumeMount := func(containerNames []string, mountName string, mountPath string, subPath string, readOnly bool) {
		if mountPath != "" {
			if containerNames == nil {
				containerNames = allContainerNames
			}
			for _, containerName := range containerNames {
				if _, ok := containerVolumeMountsMap[containerName]; !ok {
					containerVolumeMountsMap[containerName] = []*k8s.VolumeMount{}
				}
				containerVolumeMountsMap[containerName] = append(containerVolumeMountsMap[containerName], &k8s.VolumeMount{
					Name:      jsii.String(mountName),
					MountPath: jsii.String(mountPath),
					SubPath:   infrahelpers.PtrIfNonEmpty(subPath),
					ReadOnly:  infrahelpers.PtrIfNonEmpty(readOnly),
				})
			}
		}
	}
	watchTheseSecretsAndReload := []string{}
	addConfigMapHash := false
	configmapHash := sha256.New()
	for _, configmap := range props.ConfigMaps {
		volumes = append(volumes, &k8s.Volume{
			Name: jsii.String(configmap.MountName),
			ConfigMap: &k8s.ConfigMapVolumeSource{
				Name: jsii.String(configmap.Name),
			},
		})
		addVolumeMount(configmap.MountToContainers, configmap.MountName, configmap.MountPath, configmap.SubPath, configmap.ReadOnly)
		for _, key := range infrahelpers.MapKeys(configmap.Data) {
			configmapHash.Write([]byte(configmap.Data[key]))
			addConfigMapHash = true
		}
		k8s.NewKubeConfigMap(scope, jsii.String("configmap-"+configmap.Name), &k8s.KubeConfigMapProps{
			Metadata: &k8s.ObjectMeta{
				Name: jsii.String(configmap.Name),
			},
			Data: infrahelpers.PtrMap(configmap.Data),
		})
	}
	if addConfigMapHash {
		podAnnotations["configmap/checksum"] = fmt.Sprintf("%x", configmapHash.Sum(nil))
	}
	for _, secret := range props.Secrets {
		if secret.MountName != "" {
			volumes = append(volumes, &k8s.Volume{
				Name: jsii.String(secret.MountName),
				Secret: &k8s.SecretVolumeSource{
					SecretName: jsii.String(secret.Name),
				},
			})
			addVolumeMount(secret.MountToContainers, secret.MountName, secret.MountPath, secret.SubPath, secret.ReadOnly)
			if secret.MountPath != "" {
				watchTheseSecretsAndReload = append(watchTheseSecretsAndReload, secret.Name)
			}
		}
		k8s.NewKubeSecret(scope, jsii.String("secret-"+secret.Name), &k8s.KubeSecretProps{
			Metadata: &k8s.ObjectMeta{
				Name: jsii.String(secret.Name),
			},
			StringData: infrahelpers.PtrMap(secret.Data),
		})
	}
	for _, externalSecret := range props.ExternalSecrets {
		if externalSecret.MountName != "" {
			volumes = append(volumes, &k8s.Volume{
				Name: jsii.String(externalSecret.MountName),
				Secret: &k8s.SecretVolumeSource{
					SecretName: jsii.String(externalSecret.Name),
				},
			})
			addVolumeMount(externalSecret.MountToContainers, externalSecret.MountName, externalSecret.MountPath, externalSecret.SubPath, externalSecret.ReadOnly)
			if externalSecret.MountPath != "" {
				watchTheseSecretsAndReload = append(watchTheseSecretsAndReload, externalSecret.Name)
			}
		}
		NewExternalSecret(scope, jsii.String("external-secret-"+externalSecret.Name), &ExternalSecretProps{
			Name:       externalSecret.Name,
			RemoteRefs: externalSecret.RemoteRefs,
			Template:   externalSecret.Template,
		})
	}
	volumes = append(volumes, props.ExtraVolumes...)
	for _, pv := range props.PersistentVolumes {
		if pv.MountName != "" {
			volumes = append(volumes, &k8s.Volume{
				Name: jsii.String(pv.MountName),
				PersistentVolumeClaim: &k8s.PersistentVolumeClaimVolumeSource{
					ClaimName: jsii.String(pv.Name),
				},
			})
			addVolumeMount(pv.MountToContainers, pv.MountName, pv.MountPath, pv.SubPath, pv.ReadOnly)
		}
		NewPersistentVolumeClaim(scope, jsii.String("pvc-"+pv.Name), &PersistentVolumeClaim{
			Name:            pv.Name,
			StorageClass:    pv.StorageClass,
			RequestsStorage: pv.RequestsStorage,
		})
	}
	containers := []*k8s.Container{}
	servicePorts := []*k8s.ServicePort{}
	ingressHosts := []IngressHost{}
	serviceAnnotations := map[string]string{}
	for _, container := range props.Containers {
		var containerVolumeMounts []*k8s.VolumeMount
		containerVolumeMounts = append(containerVolumeMounts, containerVolumeMountsMap[container.Name]...)
		containerVolumeMounts = append(containerVolumeMounts, container.ExtraVolumeMounts...)

		env := []*k8s.EnvVar{}
		for k, v := range container.Env {
			env = append(env, &k8s.EnvVar{Name: jsii.String(k), Value: jsii.String(v)})
		}
		slices.SortFunc(env, func(a *k8s.EnvVar, b *k8s.EnvVar) bool {
			return *a.Name < *b.Name
		})
		envFrom := []*k8s.EnvFromSource{}
		for _, v := range container.EnvFromSecretRef {
			envFrom = append(envFrom, &k8s.EnvFromSource{
				SecretRef: &k8s.SecretEnvSource{
					Name: jsii.String(v),
				},
			})
			watchTheseSecretsAndReload = append(watchTheseSecretsAndReload, v)
		}
		var ports []*k8s.ContainerPort
		for _, port := range container.Ports {
			ports = append(ports, &k8s.ContainerPort{
				Name:          jsii.String(port.Name),
				ContainerPort: jsii.Number(port.Port),
			})
		}

		var args *[]*string
		if len(container.Args) > 0 {
			args = &[]*string{}
			for _, arg := range container.Args {
				*args = append(*args, jsii.String(arg))
			}
		}

		var command []*string
		if len(container.Command) > 0 {
			command = []*string{}
			for _, v := range container.Command {
				command = append(command, jsii.String(v))
			}
		}

		for _, port := range container.Ports {
			servicePorts = append(servicePorts, &k8s.ServicePort{
				Name:       jsii.String(port.Name),
				Port:       jsii.Number(port.Port),
				TargetPort: k8s.IntOrString_FromString(jsii.String(port.Name)),
			})
			if port.Ingress != nil {
				ingressHosts = append(ingressHosts, IngressHost{
					Host:  port.Ingress.Host,
					Paths: []IngressHostPath{{Path: port.Ingress.Path, ServiceName: props.Name, ServicePortName: port.Name}},
					Tls:   true,
				})
			}
			if prometheusScrape := port.PrometheusScrape; prometheusScrape != nil {
				serviceAnnotations["prometheus.io/scrape"] = "true"
				serviceAnnotations["prometheus.io/port"] = fmt.Sprint(port.Port)
				if prometheusScrape.Path != "" {
					serviceAnnotations["prometheus.io/path"] = prometheusScrape.Path
				}
			}
		}

		containers = append(containers, &k8s.Container{
			Name:    jsii.String(container.Name),
			Image:   jsii.String(container.Image.String()),
			Command: infrahelpers.PtrIfLenGt0(command),
			// ImagePullPolicy: jsii.String("IfNotPresent"),
			Env:                      infrahelpers.PtrIfLenGt0(env),
			EnvFrom:                  infrahelpers.PtrIfLenGt0(envFrom),
			Args:                     args,
			VolumeMounts:             infrahelpers.PtrIfLenGt0(containerVolumeMounts),
			Ports:                    infrahelpers.PtrIfLenGt0(ports),
			LivenessProbe:            container.LivenessProbe,
			ReadinessProbe:           container.ReadinessProbe,
			StartupProbe:             container.StartupProbe,
			TerminationMessagePolicy: jsii.String("FallbackToLogsOnError"),
			SecurityContext:          container.SecurityContext,
		})
	}
	for _, vol := range props.ExtraVolumes {
		if vol.Secret != nil {
			watchTheseSecretsAndReload = append(watchTheseSecretsAndReload, *vol.Secret.SecretName)
		}
	}
	appAnnotations := infrahelpers.CopyMap(props.AppAnnotations)
	if len(watchTheseSecretsAndReload) > 0 {
		appAnnotations["secret.reloader.stakater.com/reload"] = strings.Join(watchTheseSecretsAndReload, ",")
	}

	podTemplate := &k8s.PodTemplateSpec{
		Metadata: &k8s.ObjectMeta{
			Labels:      infrahelpers.PtrMap(commonLabels),
			Annotations: infrahelpers.PtrMap(podAnnotations),
		},
		Spec: &k8s.PodSpec{
			ServiceAccountName: infrahelpers.If(props.ServiceAccountName != "", &props.ServiceAccountName, nil),
			// AutomountServiceAccountToken: infrahelpers.If(props.ServiceAccountName != "", jsii.Bool(true), nil),
			Hostname:           infrahelpers.If(props.Hostname != "", &props.Hostname, nil),
			EnableServiceLinks: infrahelpers.If(props.EnableServiceLinks, &props.EnableServiceLinks, nil),
			SecurityContext:    props.PodSecurityContext,
			ImagePullSecrets: infrahelpers.If(props.ImagePullSecrets != "", &[]*k8s.LocalObjectReference{
				{Name: jsii.String(props.ImagePullSecrets)},
			}, nil),
			Containers:  &containers,
			Volumes:     infrahelpers.PtrIfLenGt0(volumes),
			HostNetwork: infrahelpers.If(props.HostNetwork, jsii.Bool(true), nil),
			DnsPolicy:   infrahelpers.If(props.DnsPolicy != "", &props.DnsPolicy, nil),
		},
	}
	switch props.Kind {
	case "Deployment":
		k8s.NewKubeDeployment(scope, jsii.String("deployment"), &k8s.KubeDeploymentProps{
			Metadata: &k8s.ObjectMeta{
				Name:        jsii.String(props.Name),
				Annotations: infrahelpers.PtrMap(appAnnotations),
			},
			Spec: &k8s.DeploymentSpec{
				Replicas: jsii.Number(1),
				Strategy: props.DeploymentUpdateStrategy,
				Selector: &k8s.LabelSelector{
					MatchLabels: infrahelpers.PtrMap(commonLabels),
				},
				Template: podTemplate,
			},
		})
	case "StatefulSet":
		k8s.NewKubeStatefulSet(scope, jsii.String("statefuleset"), &k8s.KubeStatefulSetProps{
			Metadata: &k8s.ObjectMeta{
				Name:        jsii.String(props.Name),
				Annotations: infrahelpers.PtrMap(appAnnotations),
			},
			Spec: &k8s.StatefulSetSpec{
				Replicas:    jsii.Number(1),
				ServiceName: jsii.String(props.Name),
				Selector: &k8s.LabelSelector{
					MatchLabels: infrahelpers.PtrMap(commonLabels),
				},
				Template: podTemplate,
			},
		})
	}
	if len(servicePorts) > 0 {
		k8s.NewKubeService(scope, jsii.String("service"), &k8s.KubeServiceProps{
			Metadata: &k8s.ObjectMeta{
				Name:        jsii.String(props.Name),
				Annotations: infrahelpers.PtrMap(serviceAnnotations),
			},
			Spec: &k8s.ServiceSpec{
				Selector: infrahelpers.PtrMap(commonLabels),
				Ports:    &servicePorts,
			},
		})
		if len(ingressHosts) > 0 {
			NewIngress(scope, jsii.String("ingress"), &IngressProps{
				Name:                   props.Name,
				Hosts:                  ingressHosts,
				TraefikMiddlewareNames: props.IngressMiddlewares,
			})
		}
	}
	return scope
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
