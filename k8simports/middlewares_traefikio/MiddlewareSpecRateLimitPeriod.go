package middlewares_traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/blesswinsamuel/infra-base/k8simports/middlewares_traefikio/jsii"
)

// Period, in combination with Average, defines the actual maximum rate, such as: r = Average / Period.
//
// It defaults to a second.
type MiddlewareSpecRateLimitPeriod interface {
	Value() interface{}
}

// The jsii proxy struct for MiddlewareSpecRateLimitPeriod
type jsiiProxy_MiddlewareSpecRateLimitPeriod struct {
	_ byte // padding
}

func (j *jsiiProxy_MiddlewareSpecRateLimitPeriod) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func MiddlewareSpecRateLimitPeriod_FromNumber(value *float64) MiddlewareSpecRateLimitPeriod {
	_init_.Initialize()

	if err := validateMiddlewareSpecRateLimitPeriod_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecRateLimitPeriod

	_jsii_.StaticInvoke(
		"middlewares_traefikio.MiddlewareSpecRateLimitPeriod",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MiddlewareSpecRateLimitPeriod_FromString(value *string) MiddlewareSpecRateLimitPeriod {
	_init_.Initialize()

	if err := validateMiddlewareSpecRateLimitPeriod_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecRateLimitPeriod

	_jsii_.StaticInvoke(
		"middlewares_traefikio.MiddlewareSpecRateLimitPeriod",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

