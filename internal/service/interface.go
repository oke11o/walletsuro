package service

//go:generate mockgen -source $GOFILE -destination mocks_interface_test.go -package ${GOPACKAGE}

import (
	"context"
	"time"

	"github.com/Rhymond/go-money"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/oke11o/walletsuro/internal/model"
)

type repo interface {
	WithTransaction(ctx context.Context, fn func(tx *sqlx.Tx) error) error

	GetWalletInTransaction(ctx context.Context, tx *sqlx.Tx, uuid uuid.UUID) (model.Wallet, error)
	GetWalletsInTransaction(ctx context.Context, tx *sqlx.Tx, fromUUID, toUUID uuid.UUID) (model.Wallet, model.Wallet, error)

	CreateWallet(ctx context.Context, tx sqlx.ExecerContext, userID int64, currency string) (model.Wallet, error)
	SaveWallet(ctx context.Context, tx *sqlx.Tx, wal model.Wallet) error
	SaveWallets(ctx context.Context, tx *sqlx.Tx, wallet model.Wallet, wallet2 model.Wallet) error

	Event(ctx context.Context, tx sqlx.ExecerContext, userID int64, amount *money.Money, walletUUID uuid.UUID, eventType string, addData *string) error

	FindEvents(ctx context.Context, userID int64, eventType *string, dateFrom, dateTo *time.Time) ([]model.Event, error)
}
