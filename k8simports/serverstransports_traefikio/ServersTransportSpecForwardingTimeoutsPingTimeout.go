package serverstransports_traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8simports/serverstransports_traefikio/jsii"
)

// PingTimeout is the timeout after which the HTTP/2 connection will be closed if a response to ping is not received.
type ServersTransportSpecForwardingTimeoutsPingTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for ServersTransportSpecForwardingTimeoutsPingTimeout
type jsiiProxy_ServersTransportSpecForwardingTimeoutsPingTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_ServersTransportSpecForwardingTimeoutsPingTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ServersTransportSpecForwardingTimeoutsPingTimeout_FromNumber(value *float64) ServersTransportSpecForwardingTimeoutsPingTimeout {
	_init_.Initialize()

	if err := validateServersTransportSpecForwardingTimeoutsPingTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportSpecForwardingTimeoutsPingTimeout

	_jsii_.StaticInvoke(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsPingTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ServersTransportSpecForwardingTimeoutsPingTimeout_FromString(value *string) ServersTransportSpecForwardingTimeoutsPingTimeout {
	_init_.Initialize()

	if err := validateServersTransportSpecForwardingTimeoutsPingTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportSpecForwardingTimeoutsPingTimeout

	_jsii_.StaticInvoke(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsPingTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

