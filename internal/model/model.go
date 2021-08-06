package model

import (
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

const DefaultCurrency = money.USD

type Wallet struct {
	UUID   uuid.UUID
	UserID int64 `db:"user_id"`
	Amount *money.Money
}

func (w *Wallet) Deposit(amount *money.Money) error {
	if amount == nil {
		return errors.New("invalid argument")
	}
	if w.Amount == nil {
		return errors.New("invalid receiver")
	}
	a, err := w.Amount.Add(amount)
	if err != nil {
		return err
	}
	w.Amount = a
	return nil
}

func (w *Wallet) Withdraw(amount *money.Money) error {
	if amount == nil {
		return errors.New("invalid argument")
	}
	if w.Amount == nil {
		return errors.New("invalid receiver")
	}
	a, err := w.Amount.Subtract(amount)
	if err != nil {
		return err
	}
	w.Amount = a
	return nil
}

func (w *Wallet) IsEnough(amount *money.Money) bool {
	if amount == nil || w.Amount == nil {
		return false
	}
	r, err := w.Amount.GreaterThanOrEqual(amount)
	if err != nil {
		return false
	}
	return r
}

type Event struct {
	UserID         int64
	Amount         *money.Money
	WalletUUID     uuid.UUID
	Type           string
	Date           time.Time
	AdditionalInfo *string
}
