package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.octolab.org/pointer"

	"github.com/oke11o/walletsuro/internal/model"
	"github.com/oke11o/walletsuro/internal/utils"
)

func New(repo repo) *Service {
	return &Service{repo: repo}
}

type Service struct {
	repo repo
}

func (s Service) CreateWallet(ctx context.Context, userID int64, currency string) (model.Wallet, error) {
	var wal model.Wallet
	err := s.repo.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		var err error
		wal, err = s.repo.CreateWallet(ctx, tx, userID, currency)
		if err != nil {
			return err
		}
		return s.repo.Event(ctx, tx, userID, money.New(0, model.DefaultCurrency), wal.UUID, model.CreateType, nil)
	})
	if err != nil {
		return wal, fmt.Errorf("cant create wallet %w", err)
	}
	return wal, nil
}

func (s Service) Deposit(ctx context.Context, userID int64, uuid uuid.UUID, amount float64) (model.Wallet, error) {
	var wal model.Wallet
	err := s.repo.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		var err error

		wal, err = s.repo.GetWalletInTransaction(ctx, tx, uuid)
		if err != nil {
			return err
		}
		depAmount := utils.FloatWithFraction(amount, wal.Amount.Currency().Fraction)
		depMoney := money.New(depAmount, wal.Amount.Currency().Code)

		if err := wal.Deposit(depMoney); err != nil {
			return err
		}
		if err := s.repo.SaveWallet(ctx, tx, wal); err != nil {
			return err
		}

		return s.repo.Event(ctx, tx, userID, depMoney, wal.UUID, model.DepositType, nil)
	})
	if err != nil {
		return model.Wallet{}, fmt.Errorf("cant create wallet %w", err)
	}

	return wal, nil
}

func (s Service) Transfer(ctx context.Context, userID int64, fromUuid uuid.UUID, toUuid uuid.UUID, amount float64) (model.Wallet, error) {
	var fromWallet model.Wallet
	err := s.repo.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		var err error
		var toWallet model.Wallet
		fromWallet, toWallet, err = s.repo.GetWalletsInTransaction(ctx, tx, fromUuid, toUuid)
		if err != nil {
			return err
		}

		transMoney := money.New(
			utils.FloatWithFraction(amount, fromWallet.Amount.Currency().Fraction),
			fromWallet.Amount.Currency().Code,
		)

		err = s.checkWalletPermission(fromWallet, userID, transMoney)
		if err != nil {
			return err
		}

		if err := fromWallet.Withdraw(transMoney); err != nil {
			return err
		}
		if err := toWallet.Deposit(transMoney); err != nil {
			return err
		}

		if err := s.repo.SaveWallets(ctx, tx, fromWallet, toWallet); err != nil {
			return err
		}

		if err := s.repo.Event(ctx, tx, userID, transMoney, fromWallet.UUID, model.WithdrawType, pointer.ToString(toWallet.UUID.String())); err != nil {
			return err
		}

		return s.repo.Event(ctx, tx, userID, transMoney, toWallet.UUID, model.DepositType, pointer.ToString(fromWallet.UUID.String()))
	})
	if err != nil {
		return model.Wallet{}, fmt.Errorf("cant create wallet %w", err)
	}

	return fromWallet, nil
}

func (s Service) Report(ctx context.Context, userID int64, eventType *string, date *time.Time) ([]model.Event, error) {
	var toDate *time.Time
	if date != nil {
		d := date.Add(24 * time.Hour)
		toDate = &d
	}
	res, err := s.repo.FindEvents(ctx, userID, eventType, date, toDate)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s Service) checkWalletPermission(wal model.Wallet, userID int64, amount *money.Money) error {
	if wal.UserID != userID {
		return model.ErrPermissionDeniedWallet
	}
	less, err := wal.Amount.LessThan(amount)
	if err != nil {
		return model.ErrPermissionDeniedWallet
	}
	if less {
		return model.ErrPermissionDeniedWallet
	}
	return nil
}
