// serverstransports_traefikcontainous
package serverstransports_traefikcontainous

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8sbase/imports/serverstransports_traefikcontainous/jsii"
)

// DialTimeout is the amount of time to wait until a connection to a backend server can be established.
type ServersTransportSpecForwardingTimeoutsDialTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for ServersTransportSpecForwardingTimeoutsDialTimeout
type jsiiProxy_ServersTransportSpecForwardingTimeoutsDialTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_ServersTransportSpecForwardingTimeoutsDialTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ServersTransportSpecForwardingTimeoutsDialTimeout_FromNumber(value *float64) ServersTransportSpecForwardingTimeoutsDialTimeout {
	_init_.Initialize()

	if err := validateServersTransportSpecForwardingTimeoutsDialTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportSpecForwardingTimeoutsDialTimeout

	_jsii_.StaticInvoke(
		"serverstransports_traefikcontainous.ServersTransportSpecForwardingTimeoutsDialTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ServersTransportSpecForwardingTimeoutsDialTimeout_FromString(value *string) ServersTransportSpecForwardingTimeoutsDialTimeout {
	_init_.Initialize()

	if err := validateServersTransportSpecForwardingTimeoutsDialTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportSpecForwardingTimeoutsDialTimeout

	_jsii_.StaticInvoke(
		"serverstransports_traefikcontainous.ServersTransportSpecForwardingTimeoutsDialTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

