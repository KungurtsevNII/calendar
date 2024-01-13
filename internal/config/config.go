package config

// Config конфигурация приложения.
type Config struct {
	Env   string           `yaml:"env" env-default:"local" validate:"required"`
	GRPC  GRPCServerConfig `yaml:"grpc_server" validate:"required"`
	Mongo MongoConfig      `yaml:"mongo_db" validate:"required"`
}
