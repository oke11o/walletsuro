package model

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Wallet struct {
	ID     int64
	UUID   uuid.UUID
	UserID int64
	amount money.Money
}
