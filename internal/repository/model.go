package repository

import (
	"time"

	"github.com/google/uuid"
)

type wallet struct {
	Uuid     uuid.UUID
	UserId   int64 `db:"user_id"`
	Amount   int64
	Currency string
}

type event struct {
	Type           string    `db:"type"`
	WalletUUID     uuid.UUID `db:"wallet_uuid"`
	Amount         int64     `db:"amount"`
	Currency       string    `db:"currency"`
	Date           time.Time `db:"date"`
	AdditionalData *string   `db:"additional_data"`
}
