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
const transferType = "transfer"

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

func (s Service) Transfer(ctx context.Context, userID int64, fromUuid uuid.UUID, toUuid uuid.UUID, amount *model.Money) (model.Wallet, error) {
	var fromWallet model.Wallet
	err := s.repo.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		var err error
		var toWallet model.Wallet
		fromWallet, toWallet, err = s.repo.GetWalletsWithBlock(ctx, tx, fromUuid, toUuid)
		if err != nil {
			return err
		}
		err = s.checkWalletPermission(fromWallet, userID, amount)
		if err != nil {
			return err
		}

		if err := fromWallet.Withdraw(amount); err != nil {
			return err
		}
		if err := toWallet.Deposit(amount); err != nil {
			return err
		}

		if err := s.repo.SaveWallets(ctx, tx, fromWallet, toWallet); err != nil {
			return err
		}

		return s.repo.Event(ctx, tx, userID, amount, fromWallet.UUID, transferType, &toWallet.UUID)
	})
	if err != nil {
		return model.Wallet{}, fmt.Errorf("cant create wallet %w", err)
	}

	return fromWallet, nil
}

func (s Service) checkWalletPermission(wal model.Wallet, userID int64, amount *model.Money) error {
	if wal.UserID != userID {
		return model.ErrPermissionDeniedWallet
	}
	less, err := wal.Amount.LessThan(&amount.Money)
	if err != nil {
		return model.ErrPermissionDeniedWallet
	}
	if less {
		return model.ErrPermissionDeniedWallet
	}
	return nil
}
