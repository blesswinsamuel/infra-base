package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PersistentVolumeClaim struct {
	Name            string
	StorageClass    string
	RequestsStorage string
	VolumeName      string
	Selector        *v1.LabelSelector
	AccessModes     []corev1.PersistentVolumeAccessMode
}

func NewPersistentVolumeClaim(scope kubegogen.Scope, props *PersistentVolumeClaim) kubegogen.ApiObject {
	return scope.AddApiObject(infrahelpers.Ptr(NewPersistentVolumeClaimProps(props)))
}

func NewPersistentVolumeClaimProps(props *PersistentVolumeClaim) corev1.PersistentVolumeClaim {
	var storageClassName *string
	if props.VolumeName != "" || props.StorageClass == "-" || props.StorageClass == "__none__" {
		storageClassName = infrahelpers.Ptr("")
	} else if props.StorageClass == "" {
		storageClassName = nil
	} else {
		storageClassName = infrahelpers.Ptr(props.StorageClass)
	}
	var resources corev1.VolumeResourceRequirements
	if props.RequestsStorage != "" {
		resources.Requests = make(corev1.ResourceList)
		resources.Requests["storage"] = resource.MustParse(props.RequestsStorage)
	}
	if props.AccessModes == nil {
		props.AccessModes = []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce}
	}
	return corev1.PersistentVolumeClaim{
		ObjectMeta: v1.ObjectMeta{
			Name: props.Name,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes:      props.AccessModes,
			Resources:        resources,
			VolumeName:       props.VolumeName,
			Selector:         props.Selector,
			StorageClassName: storageClassName,
		},
	}
}
