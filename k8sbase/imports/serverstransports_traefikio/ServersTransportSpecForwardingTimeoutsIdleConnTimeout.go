package serverstransports_traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8sbase/imports/serverstransports_traefikio/jsii"
)

// IdleConnTimeout is the maximum period for which an idle HTTP keep-alive connection will remain open before closing itself.
type ServersTransportSpecForwardingTimeoutsIdleConnTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for ServersTransportSpecForwardingTimeoutsIdleConnTimeout
type jsiiProxy_ServersTransportSpecForwardingTimeoutsIdleConnTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_ServersTransportSpecForwardingTimeoutsIdleConnTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ServersTransportSpecForwardingTimeoutsIdleConnTimeout_FromNumber(value *float64) ServersTransportSpecForwardingTimeoutsIdleConnTimeout {
	_init_.Initialize()

	if err := validateServersTransportSpecForwardingTimeoutsIdleConnTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportSpecForwardingTimeoutsIdleConnTimeout

	_jsii_.StaticInvoke(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsIdleConnTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ServersTransportSpecForwardingTimeoutsIdleConnTimeout_FromString(value *string) ServersTransportSpecForwardingTimeoutsIdleConnTimeout {
	_init_.Initialize()

	if err := validateServersTransportSpecForwardingTimeoutsIdleConnTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportSpecForwardingTimeoutsIdleConnTimeout

	_jsii_.StaticInvoke(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsIdleConnTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

