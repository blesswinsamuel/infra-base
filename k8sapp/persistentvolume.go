package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/kgen"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PersistentVolume struct {
	Name                   string
	StorageCapacity        string
	PersistentVolumeSource corev1.PersistentVolumeSource
	MountOptions           []string
	ClaimRef               *corev1.ObjectReference
	AccessModes            []corev1.PersistentVolumeAccessMode
	StorageClassName       string
	NodeAffinity           *corev1.VolumeNodeAffinity
}

func NewPersistentVolume(scope kgen.Scope, props *PersistentVolume) kgen.ApiObject {
	return scope.AddApiObject(infrahelpers.Ptr(NewPersistentVolumeProps(props)))
}

func NewPersistentVolumeProps(props *PersistentVolume) corev1.PersistentVolume {
	return corev1.PersistentVolume{
		ObjectMeta: v1.ObjectMeta{
			Name: props.Name,
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity: corev1.ResourceList{
				"storage": resource.MustParse(props.StorageCapacity),
			},
			StorageClassName:              props.StorageClassName,
			VolumeMode:                    infrahelpers.Ptr(corev1.PersistentVolumeFilesystem),
			AccessModes:                   props.AccessModes,
			PersistentVolumeReclaimPolicy: corev1.PersistentVolumeReclaimRetain,
			PersistentVolumeSource:        props.PersistentVolumeSource,
			MountOptions:                  props.MountOptions,
			ClaimRef:                      props.ClaimRef,
			NodeAffinity:                  props.NodeAffinity,
		},
	}
}
