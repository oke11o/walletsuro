package repository

import "github.com/google/uuid"

type wallet struct {
	Uuid     uuid.UUID
	UserId   int64 `db:"user_id"`
	Amount   int64
	Currency string
}
