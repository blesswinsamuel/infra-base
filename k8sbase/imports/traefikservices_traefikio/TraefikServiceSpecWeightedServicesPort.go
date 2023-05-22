package traefikservices_traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8sbase/imports/traefikservices_traefikio/jsii"
)

// Port defines the port of a Kubernetes Service.
//
// This can be a reference to a named port.
type TraefikServiceSpecWeightedServicesPort interface {
	Value() interface{}
}

// The jsii proxy struct for TraefikServiceSpecWeightedServicesPort
type jsiiProxy_TraefikServiceSpecWeightedServicesPort struct {
	_ byte // padding
}

func (j *jsiiProxy_TraefikServiceSpecWeightedServicesPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func TraefikServiceSpecWeightedServicesPort_FromNumber(value *float64) TraefikServiceSpecWeightedServicesPort {
	_init_.Initialize()

	if err := validateTraefikServiceSpecWeightedServicesPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecWeightedServicesPort

	_jsii_.StaticInvoke(
		"traefikservices_traefikio.TraefikServiceSpecWeightedServicesPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TraefikServiceSpecWeightedServicesPort_FromString(value *string) TraefikServiceSpecWeightedServicesPort {
	_init_.Initialize()

	if err := validateTraefikServiceSpecWeightedServicesPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecWeightedServicesPort

	_jsii_.StaticInvoke(
		"traefikservices_traefikio.TraefikServiceSpecWeightedServicesPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

