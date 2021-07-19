package service

//go:generate mockgen -source $GOFILE -destination mocks_interface_test.go -package ${GOPACKAGE}

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/oke11o/walletsuro/internal/model"
)

//type repo interface {
//	CreateWallet(ctx context.Context, tx sqlx.ExecerContext, userID int64) (model.Wallet, error)
//	Event(ctx context.Context, tx *sqlx.Tx, userID int64, targetWallet uuid.UUID, typ string, fromWallet *uuid.UUID) (int64, error)
//	WithTransaction(ctx context.Context, fn func(tx *sqlx.Tx) error) error
//}

type repoWalletCreater interface {
	CreateWallet(ctx context.Context, tx sqlx.ExecerContext, userID int64) (model.Wallet, error)
}
type repoEventCreater interface {
	Event(ctx context.Context, tx sqlx.ExecerContext, userID int64, targetWallet uuid.UUID, eventType string, fromWallet *uuid.UUID) error
}
type repoWithTransactioner interface {
	WithTransaction(ctx context.Context, fn func(tx *sqlx.Tx) error) error
}
