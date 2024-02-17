package config

// Config конфигурация приложения.
type Config struct {
	AppName string           `yaml:"app_name" validate:"required"`
	Env     string           `yaml:"env" env-default:"local" validate:"required"`
	GRPC    GRPCServerConfig `yaml:"grpc_server" validate:"required"`
	Mongo   MongoConfig      `yaml:"mongo_db" validate:"required"`
}
