package middlewares_traefikio


// CircuitBreaker holds the circuit breaker configuration.
type MiddlewareSpecCircuitBreaker struct {
	// CheckPeriod is the interval between successive checks of the circuit breaker condition (when in standby state).
	CheckPeriod MiddlewareSpecCircuitBreakerCheckPeriod `field:"optional" json:"checkPeriod" yaml:"checkPeriod"`
	// Expression is the condition that triggers the tripped state.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// FallbackDuration is the duration for which the circuit breaker will wait before trying to recover (from a tripped state).
	FallbackDuration MiddlewareSpecCircuitBreakerFallbackDuration `field:"optional" json:"fallbackDuration" yaml:"fallbackDuration"`
	// RecoveryDuration is the duration for which the circuit breaker will try to recover (as soon as it is in recovering state).
	RecoveryDuration MiddlewareSpecCircuitBreakerRecoveryDuration `field:"optional" json:"recoveryDuration" yaml:"recoveryDuration"`
}

