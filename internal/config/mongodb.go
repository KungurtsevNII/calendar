package config

import (
	"time"
)

// MongoConfig конфигурация драйвера MongoDB.
type MongoConfig struct {
	Endpoint               string        `yaml:"endpoint" validate:"required"`
	ConnectionTimeout      time.Duration `yaml:"connection_timeout" validate:"required"`
	ServerSelectionTimeout time.Duration `yaml:"server_selection_timeout" validate:"required"`
	MaxPoolSize            uint64        `yaml:"max_pool_size" validate:"required"`
	MaxConnIdleTime        time.Duration `yaml:"max_conn_idle_time"`
	MaxConnecting          uint64        `yaml:"max_connecting" validate:"required"`
	MinPoolSize            uint64        `yaml:"min_pool_size" validate:"required"`
}

// GetEndpoint точка подключения к MongoDB.
func (mc MongoConfig) GetEndpoint() string {
	return mc.Endpoint
}

// GetConnectionTimeout таймаут на создание соеденения с MongoDB.
// Значение по умолчанию - 30 секунд.
func (mc MongoConfig) GetConnectionTimeout() time.Duration {
	return mc.ConnectionTimeout
}

// GetServerSelectionTimeout определяет, как долго драйвер будет ждать,
// чтобы найти доступный подходящий сервер для выполнения операции.
// Значение по умолчанию - 30 секунд.
func (mc MongoConfig) GetServerSelectionTimeout() time.Duration {
	return mc.ServerSelectionTimeout
}

// GetMaxPoolSize указывает максимальное количество подключений,
// разрешенных в пуле подключений драйвера к каждому серверу.
// Если это значение равно 0, максимальный размер пула соединений не ограничен.
// Значение по умолчанию — 100.
func (mc MongoConfig) GetMaxPoolSize() uint64 {
	return mc.MaxPoolSize
}

// GetMaxConnIdleTime указывает максимальное время, в течение которого соединение
// будет оставаться бездействующим в пуле соединений прежде чем он будет удален из пула и закрыт.
// Значение 0 означает, что соединение может оставаться неиспользованным неопределенное время.
// Значение по умолчанию - 0.
func (mc MongoConfig) GetMaxConnIdleTime() time.Duration {
	return mc.ConnectionTimeout
}

// GetMaxConnecting указывает максимальное количество соединений,
// которые пул соединений может установить с сервером одновременно.
// Значение по умолчанию — 2. Значения больше 100 не рекомендуются.
func (mc MongoConfig) GetMaxConnecting() uint64 {
	return mc.MaxConnecting
}

// GetMinPoolSize указывает минимальное количество подключений,
// разрешенных в пуле подключений драйвера к каждому серверу.
// Если значение > 0, то пул каждого сервера будет поддерживаться в фоновом режиме,
// чтобы гарантировать, что размер не упадет ниже указанного значения.
// Значение по умолчанию — 0.
func (mc MongoConfig) GetMinPoolSize() uint64 {
	return mc.MinPoolSize
}
