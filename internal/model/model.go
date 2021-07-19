package model

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Wallet struct {
	UUID   uuid.UUID
	UserID int64
	amount money.Money
}

type Event struct {
	ID               int64
	UserID           int64
	TargetWalletUUID uuid.UUID
	FromWalletUUID   uuid.UUID
	Type             string
	Date             time.Time
}
