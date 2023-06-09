package middlewares_traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8simports/middlewares_traefikio/jsii"
)

// FallbackDuration is the duration for which the circuit breaker will wait before trying to recover (from a tripped state).
type MiddlewareSpecCircuitBreakerFallbackDuration interface {
	Value() interface{}
}

// The jsii proxy struct for MiddlewareSpecCircuitBreakerFallbackDuration
type jsiiProxy_MiddlewareSpecCircuitBreakerFallbackDuration struct {
	_ byte // padding
}

func (j *jsiiProxy_MiddlewareSpecCircuitBreakerFallbackDuration) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func MiddlewareSpecCircuitBreakerFallbackDuration_FromNumber(value *float64) MiddlewareSpecCircuitBreakerFallbackDuration {
	_init_.Initialize()

	if err := validateMiddlewareSpecCircuitBreakerFallbackDuration_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecCircuitBreakerFallbackDuration

	_jsii_.StaticInvoke(
		"middlewares_traefikio.MiddlewareSpecCircuitBreakerFallbackDuration",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MiddlewareSpecCircuitBreakerFallbackDuration_FromString(value *string) MiddlewareSpecCircuitBreakerFallbackDuration {
	_init_.Initialize()

	if err := validateMiddlewareSpecCircuitBreakerFallbackDuration_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecCircuitBreakerFallbackDuration

	_jsii_.StaticInvoke(
		"middlewares_traefikio.MiddlewareSpecCircuitBreakerFallbackDuration",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

