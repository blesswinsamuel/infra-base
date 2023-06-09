package k8sapp

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8simports/k8s"
)

type PersistentVolumeClaim struct {
	Name            string
	StorageClass    string
	RequestsStorage string
}

func NewPersistentVolumeClaim(scope constructs.Construct, id *string, props *PersistentVolumeClaim) k8s.KubePersistentVolumeClaim {
	return k8s.NewKubePersistentVolumeClaim(scope, id, NewPersistentVolumeClaimProps(props))
}

func NewPersistentVolumeClaimProps(props *PersistentVolumeClaim) *k8s.KubePersistentVolumeClaimProps {
	return &k8s.KubePersistentVolumeClaimProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String(props.Name),
		},
		Spec: &k8s.PersistentVolumeClaimSpec{
			AccessModes: &[]*string{jsii.String("ReadWriteOnce")},
			Resources: &k8s.ResourceRequirements{
				Requests: &map[string]k8s.Quantity{
					"storage": k8s.Quantity_FromString(&props.RequestsStorage),
				},
			},
			StorageClassName: infrahelpers.PtrIfNonEmpty(props.StorageClass),
		},
	}
}
