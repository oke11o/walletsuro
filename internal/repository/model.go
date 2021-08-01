package repository

import "github.com/google/uuid"

type wallet struct {
	Uuid     uuid.UUID
	UserId   int64
	Amount   int64
	Currency string
}
