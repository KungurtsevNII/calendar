package config

import (
	"time"
)

// GRPCServerConfig конфигурация GRPC сервера.
type GRPCServerConfig struct {
	Port    int           `yaml:"port" validate:"required,gt=0"`
	Timeout time.Duration `yaml:"timeout" validate:"required,gt=0"`
}
