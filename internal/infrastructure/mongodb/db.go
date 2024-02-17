package mongodb

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	service = "MongoDB"
)

// MongoDB обертка над стандартным драйвером Mongo.
type MongoDB struct {
	client *mongo.Client
	cfg    Config
}

// New возвращает новую обертку над драйвером MongoDB.
func New(ctx context.Context, cfg Config) (*MongoDB, error) {
	const operation = service + "New"

	opts := options.Client().
		ApplyURI(cfg.GetEndpoint()).
		SetReadPreference(readpref.SecondaryPreferred()).
		SetConnectTimeout(cfg.GetConnectionTimeout()).
		SetServerSelectionTimeout(cfg.GetServerSelectionTimeout()).
		SetMaxPoolSize(cfg.GetMaxPoolSize()).
		SetMaxConnIdleTime(cfg.GetMaxConnIdleTime()).
		SetMaxConnecting(cfg.GetMaxConnecting()).
		SetMinPoolSize(cfg.GetMinPoolSize())

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, operation)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, operation)
	}

	return &MongoDB{
		client: client,
		cfg:    cfg,
	}, nil
}

func (db *MongoDB) Disconnect(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}
