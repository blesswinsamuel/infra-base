package resourcesbase

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8s-base/imports/externalsecretsio"
	"github.com/blesswinsamuel/infra-base/k8s-base/imports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"golang.org/x/exp/slices"
)

type ApplicationProps struct {
	Kind               string                       `yaml:"kind"`
	Name               string                       `yaml:"name"`
	Annotations        map[string]string            `yaml:"annotations"`
	PodSecurityContext *k8s.PodSecurityContext      `yaml:"securityContext"`
	ImagePullSecrets   string                       `yaml:"imagePullSecrets"`
	Container          ApplicationContainer         `yaml:"container"`
	ConfigMap          *ApplicationConfigMap        `yaml:"configMap"`
	ExternalSecrets    []ApplicationExternalSecret  `yaml:"externalSecrets"`
	Secrets            []ApplicationSecret          `yaml:"secrets"`
	ExtraVolumes       []*k8s.Volume                `yaml:"extraVolumes"`
	Persistence        *ApplicationPersistence      `yaml:"persistence"`
	Ingress            []ApplicationIngress         `yaml:"ingress"`
	IngressAnnotations map[string]string            `yaml:"ingressAnnotations"`
	PrometheusScrape   *ApplicationPrometheusScrape `yaml:"prometheusScrape"`
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
	ImageInfo         ImageInfo          `yaml:"image"`
	Command           []string           `yaml:"command"`
	Env               map[string]string  `yaml:"env"`
	EnvFromSecretRef  []string           `yaml:"envFromSecretRef"`
	Args              []string           `yaml:"args"`
	Ports             []ContainerPort    `yaml:"ports"`
	ExtraVolumeMounts []*k8s.VolumeMount `yaml:"extraVolumeMounts"`
	LivenessProbe     *k8s.Probe         `yaml:"livenessProbe"`
	ReadinessProbe    *k8s.Probe         `yaml:"readinessProbe"`
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
	Name string            `yaml:"name"`
	Data map[string]string `yaml:"data"`
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
		Namespace: GetNamespace(scope),
	})
	label := map[string]*string{"app": jsii.String(props.Name)}
	var args *[]*string
	if len(props.Container.Args) > 0 {
		args = &[]*string{}
		for _, arg := range props.Container.Args {
			*args = append(*args, jsii.String(arg))
		}
	}
	var volumes []*k8s.Volume
	var volumeMounts []*k8s.VolumeMount
	var ports []*k8s.ContainerPort
	annotations := map[string]*string{}
	if props.ConfigMap != nil {
		volumes = append(volumes, &k8s.Volume{
			Name: jsii.String(props.ConfigMap.MountName),
			ConfigMap: &k8s.ConfigMapVolumeSource{
				Name: jsii.String(props.ConfigMap.Name),
			},
		})
		volumeMounts = append(volumeMounts, &k8s.VolumeMount{
			Name:      jsii.String(props.ConfigMap.MountName),
			MountPath: jsii.String(props.ConfigMap.MountPath),
			SubPath:   jsii.String(props.ConfigMap.SubPath),
			ReadOnly:  jsii.Bool(props.ConfigMap.ReadOnly),
		})
		hash := sha256.New()
		for _, v := range props.ConfigMap.Data {
			hash.Write([]byte(v))
		}
		annotations["configmap/checksum"] = jsii.String(fmt.Sprintf("%x", hash.Sum(nil)))
	}
	volumes = append(volumes, props.ExtraVolumes...)
	volumeMounts = append(volumeMounts, props.Container.ExtraVolumeMounts...)
	for k, v := range props.Annotations {
		annotations[k] = jsii.String(v)
	}
	if len(props.Container.Ports) > 0 {
		for _, port := range props.Container.Ports {
			ports = append(ports, &k8s.ContainerPort{
				Name:          jsii.String(port.Name),
				ContainerPort: jsii.Number(port.Port),
			})
		}
	}
	env := []*k8s.EnvVar{}
	for k, v := range props.Container.Env {
		env = append(env, &k8s.EnvVar{Name: jsii.String(k), Value: jsii.String(v)})
	}
	slices.SortFunc(env, func(a *k8s.EnvVar, b *k8s.EnvVar) bool {
		return *a.Name < *b.Name
	})
	envFrom := []*k8s.EnvFromSource{}
	secretReloadAnnotationValue := []string{}
	for _, v := range props.Container.EnvFromSecretRef {
		envFrom = append(envFrom, &k8s.EnvFromSource{
			SecretRef: &k8s.SecretEnvSource{
				Name: jsii.String(v),
			},
		})
		secretReloadAnnotationValue = append(secretReloadAnnotationValue, v)
	}
	topAnnotations := map[string]*string{}
	if len(secretReloadAnnotationValue) > 0 {
		topAnnotations["secret.reloader.stakater.com/reload"] = jsii.String(strings.Join(secretReloadAnnotationValue, ","))
	}
	var command []*string
	if len(props.Container.Command) > 0 {
		command = []*string{}
		for _, v := range props.Container.Command {
			command = append(command, jsii.String(v))
		}
	}
	if props.Persistence != nil {
		if props.Persistence.VolumeMountName != "" {
			volumes = append(volumes, &k8s.Volume{
				Name: jsii.String(props.Persistence.VolumeMountName),
				PersistentVolumeClaim: &k8s.PersistentVolumeClaimVolumeSource{
					ClaimName: jsii.String(props.Persistence.PersistentVolumeName),
				},
			})
			if props.Persistence.VolumeMountPath != "" {
				volumeMounts = append(volumeMounts, &k8s.VolumeMount{
					Name:      jsii.String(props.Persistence.VolumeMountName),
					MountPath: jsii.String(props.Persistence.VolumeMountPath),
				})
			}
		}
		k8s.NewKubePersistentVolumeClaim(chart, jsii.String("pvc"), &k8s.KubePersistentVolumeClaimProps{
			Metadata: &k8s.ObjectMeta{
				Name:      jsii.String(props.Persistence.PersistentVolumeName),
				Namespace: GetNamespace(scope),
			},
			Spec: &k8s.PersistentVolumeClaimSpec{
				AccessModes: &[]*string{jsii.String("ReadWriteOnce")},
				Resources: &k8s.ResourceRequirements{
					Requests: &map[string]k8s.Quantity{
						"storage": k8s.Quantity_FromString(&props.Persistence.RequestsStorage),
					},
				},
				StorageClassName: jsii.String(props.Persistence.StorageClass),
			},
		})
	}
	podTemplate := &k8s.PodTemplateSpec{
		Metadata: &k8s.ObjectMeta{
			Labels:      &label,
			Annotations: &annotations,
		},
		Spec: &k8s.PodSpec{
			SecurityContext: props.PodSecurityContext,
			ImagePullSecrets: Ternary(props.ImagePullSecrets != "", &[]*k8s.LocalObjectReference{
				{Name: jsii.String(props.ImagePullSecrets)},
			}, nil),
			Containers: &[]*k8s.Container{{
				Name:    jsii.String(props.Name),
				Image:   props.Container.ImageInfo.ToString(),
				Command: Ternary(len(command) > 0, &command, nil),
				// ImagePullPolicy: jsii.String("IfNotPresent"),
				Env:            Ternary(len(env) > 0, &env, nil),
				EnvFrom:        Ternary(len(envFrom) > 0, &envFrom, nil),
				Args:           args,
				VolumeMounts:   Ternary(len(volumeMounts) > 0, &volumeMounts, nil),
				Ports:          Ternary(len(ports) > 0, &ports, nil),
				LivenessProbe:  props.Container.LivenessProbe,
				ReadinessProbe: props.Container.ReadinessProbe,
			}},
			Volumes: &volumes,
		},
	}
	switch props.Kind {
	case "Deployment":
		k8s.NewKubeDeployment(chart, jsii.String("deployment"), &k8s.KubeDeploymentProps{
			Metadata: &k8s.ObjectMeta{
				Name:        jsii.String(props.Name),
				Namespace:   GetNamespace(scope),
				Annotations: &topAnnotations,
			},
			Spec: &k8s.DeploymentSpec{
				Replicas: jsii.Number(1),
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
				Namespace:   GetNamespace(scope),
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
				Namespace: GetNamespace(scope),
			},
			Data: data,
		})
	}
	if len(props.Container.Ports) > 0 {
		ports := &[]*k8s.ServicePort{}
		for _, port := range props.Container.Ports {
			*ports = append(*ports, &k8s.ServicePort{
				Name:       jsii.String(port.Name),
				Port:       jsii.Number(port.Port),
				TargetPort: k8s.IntOrString_FromString(jsii.String(port.Name)),
			})
		}
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
				Namespace:   GetNamespace(scope),
				Annotations: JSIIMap(serviceAnnotations),
			},
			Spec: &k8s.ServiceSpec{
				Selector: &label,
				Ports:    ports,
			},
		})
		if len(props.Ingress) > 0 {
			ingressRules := &[]*k8s.IngressRule{}
			for _, ingress := range props.Ingress {
				path := ingress.Path
				if path == nil {
					path = jsii.String("/")
				}
				*ingressRules = append(*ingressRules, &k8s.IngressRule{
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
					Namespace: GetNamespace(scope),
					Annotations: JSIIMap(MergeMaps(
						GetCertIssuerAnnotation(scope),
						props.IngressAnnotations,
					)),
				},
				Spec: &k8s.IngressSpec{
					Rules: ingressRules,
					Tls: &[]*k8s.IngressTls{
						{
							Hosts: &[]*string{
								jsii.String(props.Ingress[0].Host),
							},
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
			Template: Ternary(len(templateData) > 0, &externalsecretsio.ExternalSecretV1Beta1SpecTargetTemplate{
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
				Namespace: GetNamespace(scope),
			},
			StringData: &secrets,
		})
	}
	return chart
}
