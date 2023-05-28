package k8sbase

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/externalsecretsio"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"github.com/blesswinsamuel/infra-base/k8sbase/infraglobal"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"golang.org/x/exp/slices"
)

type ApplicationProps struct {
	Kind                     string                  `yaml:"kind"`
	Name                     string                  `yaml:"name"`
	Hostname                 string                  `yaml:"hostname"`
	DeploymentUpdateStrategy *k8s.DeploymentStrategy `yaml:"strategy"`
	TopAnnotations           map[string]string       `yaml:"topAnnotations"`
	Annotations              map[string]string       `yaml:"annotations"`
	PodSecurityContext       *k8s.PodSecurityContext `yaml:"securityContext"`
	ImagePullSecrets         string                  `yaml:"imagePullSecrets"`
	// Deprecated: use Containers instead
	Container       ApplicationContainer        `yaml:"container"`
	Containers      []ApplicationContainer      `yaml:"containers"`
	ConfigMap       *ApplicationConfigMap       `yaml:"configMap"`
	ExternalSecrets []ApplicationExternalSecret `yaml:"externalSecrets"`
	Secrets         []ApplicationSecret         `yaml:"secrets"`
	ExtraVolumes    []*k8s.Volume               `yaml:"extraVolumes"`
	// Deprecated: use PersistentVolumes instead
	Persistence        *ApplicationPersistence      `yaml:"persistence"`
	PersistentVolumes  []ApplicationPersistence     `yaml:"persistent_volumes"`
	Ingress            []ApplicationIngress         `yaml:"ingress"`
	IngressAnnotations map[string]string            `yaml:"ingressAnnotations"`
	PrometheusScrape   *ApplicationPrometheusScrape `yaml:"prometheusScrape"`
	HostNetwork        bool                         `yaml:"hostNetwork"`
	DnsPolicy          string                       `yaml:"dnsPolicy"`
}

type ApplicationIngress struct {
	Host     string  `yaml:"host"`
	Path     *string `yaml:"path"`
	PortName string  `yaml:"portName"`
}

type ApplicationPrometheusScrape struct {
	Port int    `yaml:"port"`
	Path string `yaml:"path"`
}

type ApplicationPersistence struct {
	PersistentVolumeName string `yaml:"persistentVolumeName"`
	StorageClass         string `yaml:"storageClassName"`
	RequestsStorage      string `yaml:"requestsStorage"`
	VolumeMountName      string `yaml:"volumeMountName"`
	VolumeMountPath      string `yaml:"volumeMountPath"`
}

type ApplicationContainer struct {
	Name              string             `yaml:"name"`
	ImageInfo         helpers.ImageInfo  `yaml:"image"`
	Command           []string           `yaml:"command"`
	Env               map[string]string  `yaml:"env"`
	EnvFromSecretRef  []string           `yaml:"envFromSecretRef"`
	Args              []string           `yaml:"args"`
	Ports             []ContainerPort    `yaml:"ports"`
	ExtraVolumeMounts []*k8s.VolumeMount `yaml:"extraVolumeMounts"`
	LivenessProbe     *k8s.Probe         `yaml:"livenessProbe"`
	ReadinessProbe    *k8s.Probe         `yaml:"readinessProbe"`
	StartupProbe      *k8s.Probe         `yaml:"startupProbe"`
}

type ApplicationConfigMap struct {
	Name      string            `yaml:"name"`
	Data      map[string]string `yaml:"data"`
	MountName string            `yaml:"mountName"`
	MountPath string            `yaml:"mountPath"`
	SubPath   string            `yaml:"subPath"`
	ReadOnly  bool              `yaml:"readOnly"`
}

type ApplicationExternalSecret struct {
	Name       string            `yaml:"name"`
	Template   map[string]string `yaml:"template"`
	RemoteRefs map[string]string `yaml:"remoteRefs"`
}

type ApplicationSecret struct {
	Name      string            `yaml:"name"`
	Data      map[string]string `yaml:"data"`
	MountName string            `yaml:"mountName"`
	MountPath string            `yaml:"mountPath"`
	SubPath   string            `yaml:"subPath"`
	ReadOnly  bool              `yaml:"readOnly"`
}

type ContainerPort struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
	// IngressHost string `yaml:"host"`
}

func NewApplication(scope constructs.Construct, id *string, props *ApplicationProps) cdk8s.Chart {
	if props.Kind == "" {
		props.Kind = "Deployment"
	}
	chart := cdk8s.NewChart(scope, id, &cdk8s.ChartProps{
		Namespace: helpers.GetNamespace(scope),
	})
	label := map[string]*string{"app": jsii.String(props.Name)}
	var volumes []*k8s.Volume
	annotations := map[string]*string{}
	var commonVolumeMounts []*k8s.VolumeMount
	secretReloadAnnotationValue := []string{}
	if props.ConfigMap != nil {
		volumes = append(volumes, &k8s.Volume{
			Name: jsii.String(props.ConfigMap.MountName),
			ConfigMap: &k8s.ConfigMapVolumeSource{
				Name: jsii.String(props.ConfigMap.Name),
			},
		})
		if props.ConfigMap.MountPath != "" {
			commonVolumeMounts = append(commonVolumeMounts, &k8s.VolumeMount{
				Name:      jsii.String(props.ConfigMap.MountName),
				MountPath: jsii.String(props.ConfigMap.MountPath),
				SubPath:   jsii.String(props.ConfigMap.SubPath),
				ReadOnly:  jsii.Bool(props.ConfigMap.ReadOnly),
			})
		}
		hash := sha256.New()
		for _, v := range props.ConfigMap.Data {
			hash.Write([]byte(v))
		}
		annotations["configmap/checksum"] = jsii.String(fmt.Sprintf("%x", hash.Sum(nil)))
	}
	for _, secret := range props.Secrets {
		if secret.MountName != "" {
			volumes = append(volumes, &k8s.Volume{
				Name: jsii.String(secret.MountName),
				Secret: &k8s.SecretVolumeSource{
					SecretName: jsii.String(secret.Name),
				},
			})
			if secret.MountPath != "" {
				commonVolumeMounts = append(commonVolumeMounts, &k8s.VolumeMount{
					Name:      jsii.String(secret.MountName),
					MountPath: jsii.String(secret.MountPath),
					SubPath:   jsii.String(secret.SubPath),
					ReadOnly:  jsii.Bool(secret.ReadOnly),
				})
				secretReloadAnnotationValue = append(secretReloadAnnotationValue, secret.Name)
			}
		}
	}
	volumes = append(volumes, props.ExtraVolumes...)
	for k, v := range props.Annotations {
		annotations[k] = jsii.String(v)
	}
	if props.Persistence != nil {
		props.PersistentVolumes = append(props.PersistentVolumes, *props.Persistence)
	}
	for _, pv := range props.PersistentVolumes {
		if pv.VolumeMountName != "" {
			volumes = append(volumes, &k8s.Volume{
				Name: jsii.String(pv.VolumeMountName),
				PersistentVolumeClaim: &k8s.PersistentVolumeClaimVolumeSource{
					ClaimName: jsii.String(pv.PersistentVolumeName),
				},
			})
			if pv.VolumeMountPath != "" {
				commonVolumeMounts = append(commonVolumeMounts, &k8s.VolumeMount{
					Name:      jsii.String(pv.VolumeMountName),
					MountPath: jsii.String(pv.VolumeMountPath),
				})
			}
		}
		k8s.NewKubePersistentVolumeClaim(chart, jsii.String("pvc-"+pv.PersistentVolumeName), &k8s.KubePersistentVolumeClaimProps{
			Metadata: &k8s.ObjectMeta{
				Name:      jsii.String(pv.PersistentVolumeName),
				Namespace: helpers.GetNamespace(scope),
			},
			Spec: &k8s.PersistentVolumeClaimSpec{
				AccessModes: &[]*string{jsii.String("ReadWriteOnce")},
				Resources: &k8s.ResourceRequirements{
					Requests: &map[string]k8s.Quantity{
						"storage": k8s.Quantity_FromString(&pv.RequestsStorage),
					},
				},
				StorageClassName: helpers.Ternary(
					pv.StorageClass == "",
					nil,
					jsii.String(pv.StorageClass),
				),
			},
		})
	}
	containers := []*k8s.Container{}
	if props.Container.ImageInfo.Repository != nil {
		if props.Container.Name == "" {
			props.Container.Name = props.Name
		}
		props.Containers = append(props.Containers, props.Container)
	}
	servicePorts := []*k8s.ServicePort{}
	for _, container := range props.Containers {
		var volumeMounts []*k8s.VolumeMount
		volumeMounts = append(volumeMounts, commonVolumeMounts...)
		volumeMounts = append(volumeMounts, container.ExtraVolumeMounts...)

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
			secretReloadAnnotationValue = append(secretReloadAnnotationValue, v)
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
		}

		containers = append(containers, &k8s.Container{
			Name:    jsii.String(container.Name),
			Image:   container.ImageInfo.ToString(),
			Command: helpers.Ternary(len(command) > 0, &command, nil),
			// ImagePullPolicy: jsii.String("IfNotPresent"),
			Env:                      helpers.Ternary(len(env) > 0, &env, nil),
			EnvFrom:                  helpers.Ternary(len(envFrom) > 0, &envFrom, nil),
			Args:                     args,
			VolumeMounts:             helpers.Ternary(len(volumeMounts) > 0, &volumeMounts, nil),
			Ports:                    helpers.Ternary(len(ports) > 0, &ports, nil),
			LivenessProbe:            container.LivenessProbe,
			ReadinessProbe:           container.ReadinessProbe,
			StartupProbe:             container.StartupProbe,
			TerminationMessagePolicy: jsii.String("FallbackToLogsOnError"),
		})
	}
	for _, vol := range props.ExtraVolumes {
		if vol.Secret != nil {
			secretReloadAnnotationValue = append(secretReloadAnnotationValue, *vol.Secret.SecretName)
		}
	}
	topAnnotations := map[string]*string{}
	if len(props.TopAnnotations) > 0 {
		for k, v := range props.TopAnnotations {
			topAnnotations[k] = jsii.String(v)
		}
	}
	if len(secretReloadAnnotationValue) > 0 {
		topAnnotations["secret.reloader.stakater.com/reload"] = jsii.String(strings.Join(secretReloadAnnotationValue, ","))
	}

	podTemplate := &k8s.PodTemplateSpec{
		Metadata: &k8s.ObjectMeta{
			Labels:      &label,
			Annotations: &annotations,
		},
		Spec: &k8s.PodSpec{
			Hostname:        helpers.Ternary(props.Hostname != "", &props.Hostname, nil),
			SecurityContext: props.PodSecurityContext,
			ImagePullSecrets: helpers.Ternary(props.ImagePullSecrets != "", &[]*k8s.LocalObjectReference{
				{Name: jsii.String(props.ImagePullSecrets)},
			}, nil),
			Containers:  &containers,
			Volumes:     &volumes,
			HostNetwork: helpers.Ternary(props.HostNetwork, jsii.Bool(true), nil),
			DnsPolicy:   helpers.Ternary(props.DnsPolicy != "", &props.DnsPolicy, nil),
		},
	}
	switch props.Kind {
	case "Deployment":
		k8s.NewKubeDeployment(chart, jsii.String("deployment"), &k8s.KubeDeploymentProps{
			Metadata: &k8s.ObjectMeta{
				Name:        jsii.String(props.Name),
				Namespace:   helpers.GetNamespace(scope),
				Annotations: &topAnnotations,
			},
			Spec: &k8s.DeploymentSpec{
				Replicas: jsii.Number(1),
				Strategy: props.DeploymentUpdateStrategy,
				Selector: &k8s.LabelSelector{
					MatchLabels: &label,
				},
				Template: podTemplate,
			},
		})
	case "StatefulSet":
		k8s.NewKubeStatefulSet(chart, jsii.String("statefuleset"), &k8s.KubeStatefulSetProps{
			Metadata: &k8s.ObjectMeta{
				Name:        jsii.String(props.Name),
				Namespace:   helpers.GetNamespace(scope),
				Annotations: &topAnnotations,
			},
			Spec: &k8s.StatefulSetSpec{
				Replicas:    jsii.Number(1),
				ServiceName: jsii.String(props.Name),
				Selector: &k8s.LabelSelector{
					MatchLabels: &label,
				},
				Template: podTemplate,
			},
		})
	}
	if props.ConfigMap != nil {
		data := &map[string]*string{}
		for k, v := range props.ConfigMap.Data {
			(*data)[k] = jsii.String(v)
		}
		k8s.NewKubeConfigMap(chart, jsii.String("configmap"), &k8s.KubeConfigMapProps{
			Metadata: &k8s.ObjectMeta{
				Name:      jsii.String(props.ConfigMap.Name),
				Namespace: helpers.GetNamespace(scope),
			},
			Data: data,
		})
	}
	if len(servicePorts) > 0 {
		serviceAnnotations := map[string]string{}
		if props.PrometheusScrape != nil {
			serviceAnnotations["prometheus.io/scrape"] = "true"
			serviceAnnotations["prometheus.io/port"] = fmt.Sprint(props.PrometheusScrape.Port)
			if props.PrometheusScrape.Path != "" {
				serviceAnnotations["prometheus.io/path"] = props.PrometheusScrape.Path
			}
		}
		k8s.NewKubeService(chart, jsii.String("service"), &k8s.KubeServiceProps{
			Metadata: &k8s.ObjectMeta{
				Name:        jsii.String(props.Name),
				Namespace:   helpers.GetNamespace(scope),
				Annotations: helpers.JSIIMap(serviceAnnotations),
			},
			Spec: &k8s.ServiceSpec{
				Selector: &label,
				Ports:    &servicePorts,
			},
		})
		if len(props.Ingress) > 0 {
			ingressRules := []*k8s.IngressRule{}
			tlsHosts := map[string]bool{}
			for _, ingress := range props.Ingress {
				path := ingress.Path
				if path == nil {
					path = jsii.String("/")
				}
				tlsHosts[ingress.Host] = true
				ingressRules = append(ingressRules, &k8s.IngressRule{
					Host: jsii.String(ingress.Host),
					Http: &k8s.HttpIngressRuleValue{
						Paths: &[]*k8s.HttpIngressPath{
							{
								Path:     path,
								PathType: jsii.String("Prefix"),
								Backend: &k8s.IngressBackend{
									Service: &k8s.IngressServiceBackend{
										Name: jsii.String(props.Name),
										Port: &k8s.ServiceBackendPort{
											Name: jsii.String(ingress.PortName),
										},
									},
								},
							},
						},
					},
				})
			}
			k8s.NewKubeIngress(chart, jsii.String("ingress"), &k8s.KubeIngressProps{
				Metadata: &k8s.ObjectMeta{
					Name:      jsii.String(props.Name),
					Namespace: helpers.GetNamespace(scope),
					Annotations: helpers.JSIIMap(helpers.MergeMaps(
						infraglobal.GetCertIssuerAnnotation(scope),
						props.IngressAnnotations,
					)),
				},
				Spec: &k8s.IngressSpec{
					Rules: &ingressRules,
					Tls: &[]*k8s.IngressTls{
						{
							Hosts:      helpers.JSIISlice(helpers.MapKeys(tlsHosts)...),
							SecretName: jsii.String(fmt.Sprintf("%s-tls", props.Name)),
						},
					},
				},
			})
		}
	}
	for _, externalSecret := range props.ExternalSecrets {
		templateData := map[string]*string{}
		for k, v := range externalSecret.Template {
			templateData[k] = jsii.String(v)
		}
		NewExternalSecret(chart, jsii.String("external-secret-"+externalSecret.Name), &ExternalSecretProps{
			Name:            jsii.String(externalSecret.Name),
			RefreshInterval: jsii.String("10m"),
			Secrets:         externalSecret.RemoteRefs,
			Template: helpers.Ternary(len(templateData) > 0, &externalsecretsio.ExternalSecretV1Beta1SpecTargetTemplate{
				Data: &templateData,
			}, nil),
		})
	}
	for _, secret := range props.Secrets {
		secrets := map[string]*string{}
		for k, v := range secret.Data {
			secrets[k] = jsii.String(v)
		}
		k8s.NewKubeSecret(chart, jsii.String("secret-"+secret.Name), &k8s.KubeSecretProps{
			Metadata: &k8s.ObjectMeta{
				Name:      jsii.String(secret.Name),
				Namespace: helpers.GetNamespace(scope),
			},
			StringData: &secrets,
		})
	}
	return chart
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
