package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
	"github.com/oke11o/walletsuro/internal/model"
)

type eventType string

const createType = "create"
const depositType = "deposit"

func New(repo repo) *Service {
	return &Service{repo: repo}
}

type Service struct {
	repo repo
}

func (s Service) CreateWallet(ctx context.Context, userID int64) (model.Wallet, error) {
	var wal model.Wallet
	err := s.repo.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		var err error
		wal, err = s.repo.CreateWallet(ctx, tx, userID)
		if err != nil {
			return err
		}
		return s.repo.Event(ctx, tx, userID, model.NewMoney(0, model.DefaultCurrency), wal.UUID, createType, nil)
	})
	if err != nil {
		return wal, fmt.Errorf("cant create wallet %w", err)
	}
	return wal, nil
}

func (s Service) Deposit(ctx context.Context, userID int64, uuid uuid.UUID, amount *model.Money) (model.Wallet, error) {
	var wal model.Wallet
	err := s.repo.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		var err error

		wal, err = s.repo.GetWalletWithBlock(ctx, tx, uuid)
		if err != nil {
			return err
		}

		if err := wal.Deposit(amount); err != nil {
			return err
		}
		if err := s.repo.SaveWallet(ctx, tx, wal); err != nil {
			return err
		}

		return s.repo.Event(ctx, tx, userID, amount, wal.UUID, depositType, nil)
	})
	if err != nil {
		return model.Wallet{}, fmt.Errorf("cant create wallet %w", err)
	}

	return wal, nil
}
