package use_cases

import (
	"context"

	"github.com/google/uuid"

	"github.com/KungurtsevNII/calendar/internal/domain"
)

func (s *Service) GetUserByID(ctx context.Context, userID uuid.UUID) (domain.User, error) {
	return s.userRepo.GetUserByID(ctx, userID)
}
