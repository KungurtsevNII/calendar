package mongodb

import (
	"time"

	"github.com/google/uuid"

	"github.com/KungurtsevNII/calendar/internal/domain"
)

type UserRecord struct {
	UserID      string    `bson:"user_id"`
	Email       string    `bson:"email"`
	FirstName   string    `bson:"first_name"`
	LastName    string    `bson:"last_name"`
	Patronymic  string    `bson:"patronymic"`
	DateOfBirth time.Time `bson:"date_of_birth"`
}

func (us UserRecord) ToDomain() (*domain.User, error) {
	userUUID, err := uuid.Parse(us.UserID)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID: userUUID,
	}, nil
}
