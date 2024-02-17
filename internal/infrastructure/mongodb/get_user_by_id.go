package mongodb

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/KungurtsevNII/calendar/internal/domain"
)

const (
	getUserByIDTimeout = 200 * time.Millisecond
)

func (db *MongoDB) GetUserByID(ctx context.Context, userID uuid.UUID) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), getUserByIDTimeout)
	defer cancel()

	return domain.User{}, nil
}
