package serverstransports_traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8simports/serverstransports_traefikio/jsii"
)

// ReadIdleTimeout is the timeout after which a health check using ping frame will be carried out if no frame is received on the HTTP/2 connection.
type ServersTransportSpecForwardingTimeoutsReadIdleTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for ServersTransportSpecForwardingTimeoutsReadIdleTimeout
type jsiiProxy_ServersTransportSpecForwardingTimeoutsReadIdleTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_ServersTransportSpecForwardingTimeoutsReadIdleTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ServersTransportSpecForwardingTimeoutsReadIdleTimeout_FromNumber(value *float64) ServersTransportSpecForwardingTimeoutsReadIdleTimeout {
	_init_.Initialize()

	if err := validateServersTransportSpecForwardingTimeoutsReadIdleTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportSpecForwardingTimeoutsReadIdleTimeout

	_jsii_.StaticInvoke(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsReadIdleTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ServersTransportSpecForwardingTimeoutsReadIdleTimeout_FromString(value *string) ServersTransportSpecForwardingTimeoutsReadIdleTimeout {
	_init_.Initialize()

	if err := validateServersTransportSpecForwardingTimeoutsReadIdleTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportSpecForwardingTimeoutsReadIdleTimeout

	_jsii_.StaticInvoke(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsReadIdleTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

