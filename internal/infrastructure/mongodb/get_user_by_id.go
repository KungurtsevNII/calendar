package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/KungurtsevNII/calendar/internal/domain"
)

const (
	getUserByIDTimeout = 200000 * time.Millisecond
)

func (db *MongoDB) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), getUserByIDTimeout)
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
