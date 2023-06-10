package k8sapp

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	externalsecretsv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var scheme = runtime.NewScheme()

func init() {
	if err := corev1.AddToScheme(scheme); err != nil {
		panic(err)
	}
	if err := externalsecretsv1beta1.AddToScheme(scheme); err != nil {
		panic(err)
	}
	if err := certmanagerv1.AddToScheme(scheme); err != nil {
		panic(err)
	}
}

func NewK8sObject(scope constructs.Construct, id *string, obj runtime.Object) cdk8s.ApiObject {
	groupVersionKinds, _, err := scheme.ObjectKinds(obj)
	if err != nil {
		panic(err)
	}
	if len(groupVersionKinds) != 1 {
		panic(fmt.Errorf("expected 1 groupVersionKind, got %d: %v", len(groupVersionKinds), groupVersionKinds))
	}
	var metadata *cdk8s.ApiObjectMetadata
	if obj, ok := obj.(metav1.Object); ok {
		if obj.GetNamespace() == "" {
			if namespaceCtx := GetNamespaceContext(scope); namespaceCtx != "" {
				obj.SetNamespace(namespaceCtx)
			}
		}
		metadata = &cdk8s.ApiObjectMetadata{Name: jsii.String(obj.GetName()), Namespace: infrahelpers.PtrIfNonEmpty(obj.GetNamespace())}
	}
	groupVersion := groupVersionKinds[0]
	apiobj := cdk8s.NewApiObject(scope, id, &cdk8s.ApiObjectProps{
		ApiVersion: jsii.String(groupVersion.GroupVersion().String()),
		Kind:       jsii.String(groupVersion.Kind),
		Metadata:   metadata,
	})
	apiobj.AddJsonPatch(cdk8s.JsonPatch_Replace(jsii.String("/spec"), infrahelpers.K8sObjectToMap(obj)["spec"]))
	// apiobj.AddJsonPatch(cdk8s.JsonPatch_Replace(jsii.String("/metadata"), infrahelpers.K8sObjectToMap(obj)["metadata"]))
	return apiobj
}
