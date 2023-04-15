// middlewares_traefikcontainous
package middlewares_traefikcontainous

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8sbase/imports/middlewares_traefikcontainous/jsii"
)

// CheckPeriod is the interval between successive checks of the circuit breaker condition (when in standby state).
type MiddlewareSpecCircuitBreakerCheckPeriod interface {
	Value() interface{}
}

// The jsii proxy struct for MiddlewareSpecCircuitBreakerCheckPeriod
type jsiiProxy_MiddlewareSpecCircuitBreakerCheckPeriod struct {
	_ byte // padding
}

func (j *jsiiProxy_MiddlewareSpecCircuitBreakerCheckPeriod) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func MiddlewareSpecCircuitBreakerCheckPeriod_FromNumber(value *float64) MiddlewareSpecCircuitBreakerCheckPeriod {
	_init_.Initialize()

	if err := validateMiddlewareSpecCircuitBreakerCheckPeriod_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecCircuitBreakerCheckPeriod

	_jsii_.StaticInvoke(
		"middlewares_traefikcontainous.MiddlewareSpecCircuitBreakerCheckPeriod",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MiddlewareSpecCircuitBreakerCheckPeriod_FromString(value *string) MiddlewareSpecCircuitBreakerCheckPeriod {
	_init_.Initialize()

	if err := validateMiddlewareSpecCircuitBreakerCheckPeriod_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecCircuitBreakerCheckPeriod

	_jsii_.StaticInvoke(
		"middlewares_traefikcontainous.MiddlewareSpecCircuitBreakerCheckPeriod",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

