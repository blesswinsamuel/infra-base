package middlewares_traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8sbase/imports/middlewares_traefikio/jsii"
)

// InitialInterval defines the first wait time in the exponential backoff series.
//
// The maximum interval is calculated as twice the initialInterval. If unspecified, requests will be retried immediately. The value of initialInterval should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.
type MiddlewareSpecRetryInitialInterval interface {
	Value() interface{}
}

// The jsii proxy struct for MiddlewareSpecRetryInitialInterval
type jsiiProxy_MiddlewareSpecRetryInitialInterval struct {
	_ byte // padding
}

func (j *jsiiProxy_MiddlewareSpecRetryInitialInterval) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func MiddlewareSpecRetryInitialInterval_FromNumber(value *float64) MiddlewareSpecRetryInitialInterval {
	_init_.Initialize()

	if err := validateMiddlewareSpecRetryInitialInterval_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecRetryInitialInterval

	_jsii_.StaticInvoke(
		"middlewares_traefikio.MiddlewareSpecRetryInitialInterval",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MiddlewareSpecRetryInitialInterval_FromString(value *string) MiddlewareSpecRetryInitialInterval {
	_init_.Initialize()

	if err := validateMiddlewareSpecRetryInitialInterval_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecRetryInitialInterval

	_jsii_.StaticInvoke(
		"middlewares_traefikio.MiddlewareSpecRetryInitialInterval",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

