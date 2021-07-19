package service

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/oke11o/walletsuro/internal/model"
)

type eventType string

const createType = "create"

func New(repoWalletCreater repoWalletCreater, repoEventCreater repoEventCreater, repoWithTransactioner repoWithTransactioner) *Service {
	return &Service{
		repoWalletCreater:     repoWalletCreater,
		repoEventCreater:      repoEventCreater,
		repoWithTransactioner: repoWithTransactioner,
	}
}

type Service struct {
	//repo repo

	repoWalletCreater     repoWalletCreater
	repoEventCreater      repoEventCreater
	repoWithTransactioner repoWithTransactioner
}

func (s Service) CreateWallet(ctx context.Context, userID int64) (model.Wallet, error) {
	var wal model.Wallet
	err := s.repoWithTransactioner.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		var err error
		wal, err = s.repoWalletCreater.CreateWallet(ctx, tx, userID)
		if err != nil {
			return err
		}
		return s.repoEventCreater.Event(ctx, tx, userID, wal.UUID, createType, nil)
	})
	if err != nil {
		return wal, fmt.Errorf("cant create wallet %w", err)
	}
	return wal, nil
}
