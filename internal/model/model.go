package model

import (
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

const DefaultCurrency = money.USD

type Wallet struct {
	UUID   uuid.UUID
	UserID int64 `db:"user_id"`
	Amount *Money
}

func (w *Wallet) Deposit(amount *Money) error {
	a, err := w.Amount.Add(&amount.Money)
	if err != nil {
		return err
	}
	w.Amount = &Money{Money: *a}
	return nil
}

func (w *Wallet) Withdraw(amount *Money) error {
	a, err := w.Amount.Subtract(&amount.Money)
	if err != nil {
		return err
	}
	w.Amount = &Money{Money: *a}
	return nil
}

func (w *Wallet) IsEnough(amount *Money) bool {
	r, err := w.Amount.GreaterThanOrEqual(&amount.Money)
	if err != nil {
		return false
	}
	return r
}

func NewMoney(amount int64, currency string) *Money {
	m := money.New(amount, currency)
	if m == nil {
		return nil
	}
	return &Money{Money: *m}
}

type Money struct {
	money.Money
}

func (m *Money) Scan(raw interface{}) error {
	switch v := raw.(type) {
	case int64:
		mon := money.New(v, DefaultCurrency)
		*m = Money{*mon}
	case int32:
		mon := money.New(int64(v), DefaultCurrency)
		*m = Money{*mon}
	case int:
		mon := money.New(int64(v), DefaultCurrency)
		*m = Money{*mon}
	case nil:
		return nil
	default:
		return fmt.Errorf("cannot sql.Scan() Money from: %#v", v)
	}
	return nil
}

type Event struct {
	ID               int64
	UserID           int64
	TargetWalletUUID uuid.UUID
	FromWalletUUID   uuid.UUID
	Type             string
	Date             time.Time
}
