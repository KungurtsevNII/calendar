package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Email       string
	FirstName   string
	LastName    string
	Patronymic  string
	DateOfBirth time.Time
}
