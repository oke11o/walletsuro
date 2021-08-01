package service

//go:generate mockgen -source $GOFILE -destination mocks_interface_test.go -package ${GOPACKAGE}

import (
	"context"
	"github.com/Rhymond/go-money"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/oke11o/walletsuro/internal/model"
)

type repo interface {
	CreateWallet(ctx context.Context, tx sqlx.ExecerContext, userID int64) (model.Wallet, error)
	Event(ctx context.Context, tx sqlx.ExecerContext, userID int64, amount *money.Money, targetWallet uuid.UUID, eventType string, fromWallet *uuid.UUID) error
	WithTransaction(ctx context.Context, fn func(tx *sqlx.Tx) error) error
	GetWalletWithBlock(ctx context.Context, tx *sqlx.Tx, uuid uuid.UUID) (model.Wallet, error)
	GetWalletsWithBlock(ctx context.Context, tx *sqlx.Tx, fromUUID, toUUID uuid.UUID) (model.Wallet, model.Wallet, error)
	SaveWallet(ctx context.Context, tx *sqlx.Tx, wal model.Wallet) error
	SaveWallets(ctx context.Context, tx *sqlx.Tx, wallet model.Wallet, wallet2 model.Wallet) error
	FindEvents(ctx context.Context, id int64, t *string, date *time.Time) ([]model.Event, error)
}
