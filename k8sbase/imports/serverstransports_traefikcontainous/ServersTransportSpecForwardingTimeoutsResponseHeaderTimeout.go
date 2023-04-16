// serverstransports_traefikcontainous
package serverstransports_traefikcontainous

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8sbase/imports/serverstransports_traefikcontainous/jsii"
)

// ResponseHeaderTimeout is the amount of time to wait for a server's response headers after fully writing the request (including its body, if any).
type ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout
type jsiiProxy_ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout_FromNumber(value *float64) ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout {
	_init_.Initialize()

	if err := validateServersTransportSpecForwardingTimeoutsResponseHeaderTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout

	_jsii_.StaticInvoke(
		"serverstransports_traefikcontainous.ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout_FromString(value *string) ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout {
	_init_.Initialize()

	if err := validateServersTransportSpecForwardingTimeoutsResponseHeaderTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout

	_jsii_.StaticInvoke(
		"serverstransports_traefikcontainous.ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

