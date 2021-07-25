package handler

//go:generate mockgen -source $GOFILE -destination mocks_interface_test.go -package ${GOPACKAGE}

import (
	"context"

	"github.com/google/uuid"

	"github.com/oke11o/walletsuro/internal/model"
)

type service interface {
	CreateWallet(ctx context.Context, userID int64) (model.Wallet, error)
	Deposit(ctx context.Context, userID int64, uuid uuid.UUID, amount *model.Money) (model.Wallet, error)
	Transfer(ctx context.Context, userID int64, fromWalletUUID uuid.UUID, toWalletUUID uuid.UUID, amount *model.Money) (model.Wallet, error)
}
