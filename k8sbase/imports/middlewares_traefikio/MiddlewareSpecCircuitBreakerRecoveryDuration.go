package middlewares_traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8sbase/imports/middlewares_traefikio/jsii"
)

// RecoveryDuration is the duration for which the circuit breaker will try to recover (as soon as it is in recovering state).
type MiddlewareSpecCircuitBreakerRecoveryDuration interface {
	Value() interface{}
}

// The jsii proxy struct for MiddlewareSpecCircuitBreakerRecoveryDuration
type jsiiProxy_MiddlewareSpecCircuitBreakerRecoveryDuration struct {
	_ byte // padding
}

func (j *jsiiProxy_MiddlewareSpecCircuitBreakerRecoveryDuration) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func MiddlewareSpecCircuitBreakerRecoveryDuration_FromNumber(value *float64) MiddlewareSpecCircuitBreakerRecoveryDuration {
	_init_.Initialize()

	if err := validateMiddlewareSpecCircuitBreakerRecoveryDuration_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecCircuitBreakerRecoveryDuration

	_jsii_.StaticInvoke(
		"middlewares_traefikio.MiddlewareSpecCircuitBreakerRecoveryDuration",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MiddlewareSpecCircuitBreakerRecoveryDuration_FromString(value *string) MiddlewareSpecCircuitBreakerRecoveryDuration {
	_init_.Initialize()

	if err := validateMiddlewareSpecCircuitBreakerRecoveryDuration_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecCircuitBreakerRecoveryDuration

	_jsii_.StaticInvoke(
		"middlewares_traefikio.MiddlewareSpecCircuitBreakerRecoveryDuration",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

