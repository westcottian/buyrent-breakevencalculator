package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

const (
	envDevelopment = "development"
	envProduction  = "production"
	envTest        = "test"
	envLocal       = "local"
)

// Env stores configuration settings extract from environmental variables
// by using https://github.com/kelseyhightower/envconfig
//
// The practice getting from environmental variables comes from https://12factor.net.
type Env struct {

	// Env is environment where application is running. This value is used to
	// annotate datadog metrics or sentry error reporting. The value must be
	// "development" or "production".
	Env string `envconfig:"ENV" required:"true"`

	// GRPCPort is port to listen. Default is 5000.
	GRPCPort int `envconfig:"GRPC_PORT" default:"5000"`
}

// IsProduction returns true if it is production environment.
func (e *Env) IsProduction() bool {
	return e.Env == envProduction
}

// IsLocal returns true if it is local environment.
func (e *Env) IsLocal() bool {
	return e.Env == envLocal
}

// ReadFromEnv reads configuration from environmental variables
// defined by Env struct.
func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Wrap(err, "failed to process envconfig")
	}

	return &env, nil
}
