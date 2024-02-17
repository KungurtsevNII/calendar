package use_cases

import (
	"context"

	"github.com/google/uuid"

	"github.com/KungurtsevNII/calendar/internal/domain"
)

type userRepo interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (domain.User, error)
}

type Service struct {
	userRepo userRepo
}

func New(userRepo userRepo) (*Service, error) {
	return &Service{
		userRepo: userRepo,
	}, nil
}
