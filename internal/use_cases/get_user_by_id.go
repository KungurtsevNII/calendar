package use_cases

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/KungurtsevNII/calendar/internal/domain"
)

func (s *Service) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	if userID == "" {
		return nil, errors.New("asdasd")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("asdasd")
	}

	return s.userRepo.GetUserByID(ctx, userUUID)
}
