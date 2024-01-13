package mongodb

import (
	"time"
)

type Config interface {
	GetEndpoint() string
	GetConnectionTimeout() time.Duration
	GetServerSelectionTimeout() time.Duration
	GetMaxPoolSize() uint64
	GetMaxConnIdleTime() time.Duration
	GetMaxConnecting() uint64
	GetMinPoolSize() uint64
}
