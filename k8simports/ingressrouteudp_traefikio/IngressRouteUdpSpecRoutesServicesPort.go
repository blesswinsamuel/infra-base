package ingressrouteudp_traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8simports/ingressrouteudp_traefikio/jsii"
)

// Port defines the port of a Kubernetes Service.
//
// This can be a reference to a named port.
type IngressRouteUdpSpecRoutesServicesPort interface {
	Value() interface{}
}

// The jsii proxy struct for IngressRouteUdpSpecRoutesServicesPort
type jsiiProxy_IngressRouteUdpSpecRoutesServicesPort struct {
	_ byte // padding
}

func (j *jsiiProxy_IngressRouteUdpSpecRoutesServicesPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func IngressRouteUdpSpecRoutesServicesPort_FromNumber(value *float64) IngressRouteUdpSpecRoutesServicesPort {
	_init_.Initialize()

	if err := validateIngressRouteUdpSpecRoutesServicesPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns IngressRouteUdpSpecRoutesServicesPort

	_jsii_.StaticInvoke(
		"ingressrouteudp_traefikio.IngressRouteUdpSpecRoutesServicesPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func IngressRouteUdpSpecRoutesServicesPort_FromString(value *string) IngressRouteUdpSpecRoutesServicesPort {
	_init_.Initialize()

	if err := validateIngressRouteUdpSpecRoutesServicesPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns IngressRouteUdpSpecRoutesServicesPort

	_jsii_.StaticInvoke(
		"ingressrouteudp_traefikio.IngressRouteUdpSpecRoutesServicesPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

