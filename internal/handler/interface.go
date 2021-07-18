package handler

//go:generate mockgen -source $GOFILE -destination mocks_interface_test.go -package ${GOPACKAGE}

import (
	"context"

	"github.com/oke11o/walletsuro/internal/model"
)

type service interface {
	CreateWallet(ctx context.Context, userID int64) (model.Wallet, error)
}
