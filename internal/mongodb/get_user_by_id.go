package mongodb

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/KungurtsevNII/calendar/internal/domain"
)

func (db *MongoDB) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.cfg.GetUserByIDTimeout())
	defer cancel()

	var record UserRecord
	filter := bson.M{"user_id": userID.String()}
	err := db.userCollection.FindOne(ctx, filter).Decode(&record)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, err
	case err != nil:
		return nil, err
	}

	return record.ToDomain()
}
