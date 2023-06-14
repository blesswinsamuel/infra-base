package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/packager"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PersistentVolumeClaim struct {
	Name            string
	StorageClass    string
	RequestsStorage string
}

func NewPersistentVolumeClaim(scope packager.Construct, id string, props *PersistentVolumeClaim) packager.ApiObject {
	return NewK8sObject(scope, id, infrahelpers.Ptr(NewPersistentVolumeClaimProps(props)))
}

func NewPersistentVolumeClaimProps(props *PersistentVolumeClaim) corev1.PersistentVolumeClaim {
	return corev1.PersistentVolumeClaim{
		ObjectMeta: v1.ObjectMeta{
			Name: props.Name,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{"ReadWriteOnce"},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					"storage": resource.MustParse(props.RequestsStorage),
				},
			},
			StorageClassName: infrahelpers.PtrIfNonEmpty(props.StorageClass),
		},
	}
}
