package traefikservices_traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8simports/traefikservices_traefikio/jsii"
)

// Port defines the port of a Kubernetes Service.
//
// This can be a reference to a named port.
type TraefikServiceSpecMirroringPort interface {
	Value() interface{}
}

// The jsii proxy struct for TraefikServiceSpecMirroringPort
type jsiiProxy_TraefikServiceSpecMirroringPort struct {
	_ byte // padding
}

func (j *jsiiProxy_TraefikServiceSpecMirroringPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func TraefikServiceSpecMirroringPort_FromNumber(value *float64) TraefikServiceSpecMirroringPort {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringPort

	_jsii_.StaticInvoke(
		"traefikservices_traefikio.TraefikServiceSpecMirroringPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TraefikServiceSpecMirroringPort_FromString(value *string) TraefikServiceSpecMirroringPort {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringPort

	_jsii_.StaticInvoke(
		"traefikservices_traefikio.TraefikServiceSpecMirroringPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

