package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/oke11o/walletsuro/internal/config"
	"github.com/oke11o/walletsuro/internal/model"
)

func New(cfg config.Config) (*Repo, error) {
	db, err := sqlx.Connect("postgres", cfg.PgDSN)
	if err != nil {
		return nil, fmt.Errorf("postrgess connect err: %w", err)
	}
	return &Repo{db: db}, nil
}

type Repo struct {
	db *sqlx.DB
}

func (Repo) CreateWallet(ctx context.Context, tx sqlx.ExecerContext, userID int64) (model.Wallet, error) {
	UUID := uuid.New()
	sql := "INSERT INTO wallets (user_id, uuid) VALUES ($1, $2)"
	_, err := tx.ExecContext(ctx, sql, userID, UUID.String())
	if err != nil {
		return model.Wallet{}, fmt.Errorf("cant create wallet: %w", err)
	}

	wallet := model.Wallet{
		UUID:   UUID,
		UserID: userID,
	}
	return wallet, nil
}

func (Repo) Event(ctx context.Context, tx sqlx.ExecerContext, userID int64, amount *model.Money, targetWallet uuid.UUID, eventType string, fromWallet *uuid.UUID) error {
	args := make([]interface{}, 0, 3)
	args = append(args, userID, amount.Amount(), targetWallet.String(), eventType)
	var sql string
	if fromWallet != nil {
		args = append(args, fromWallet.String())
		sql = "INSERT INTO events (user_id, amount, target_wallet_uuid, type, from_wallet_uuid) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	} else {
		sql = "INSERT INTO events (user_id, amount, target_wallet_uuid, type) VALUES ($1, $2, $3, $4) RETURNING id"
	}

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("cant save event: %w", err)
	}
	return nil
}

func (r *Repo) WithTransaction(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	var err error
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("beginning tx: %w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Printf("Failed to rollback transaction: %v", err)
		}
	}()

	err = fn(tx)
	if err != nil {
		return fmt.Errorf("executing transactional func: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("committing tx: %w", err)
	}
	return nil
}

func (r *Repo) GetWalletWithBlock(ctx context.Context, tx *sqlx.Tx, uuid uuid.UUID) (model.Wallet, error) {
	var val model.Wallet
	s := "SELECT uuid, user_id, amount FROM wallets WHERE uuid=$1 FOR UPDATE SKIP LOCKED"
	if err := tx.GetContext(ctx, &val, s, uuid); err != nil {
		return val, err
	}

	return val, nil
}

func (r *Repo) GetWalletsWithBlock(ctx context.Context, tx *sqlx.Tx, fromUUID, toUUID uuid.UUID) (model.Wallet, model.Wallet, error) {
	var val []model.Wallet
	s := "SELECT uuid, user_id, amount FROM wallets WHERE uuid=$1 OR uuid=$2 FOR UPDATE SKIP LOCKED"
	if err := tx.SelectContext(ctx, &val, s, fromUUID, toUUID); err != nil {
		return model.Wallet{}, model.Wallet{}, err
	}
	if len(val) != 2 {
		return model.Wallet{}, model.Wallet{}, errors.New("cant get both wallets")
	}
	var from, to model.Wallet
	if fromUUID == val[0].UUID {
		from = val[0]
		to = val[1]
	} else {
		to = val[0]
		from = val[1]
	}

	return from, to, nil
}

func (r *Repo) SaveWallet(ctx context.Context, tx *sqlx.Tx, wal model.Wallet) error {
	s := "INSERT INTO wallets (uuid, user_id, amount) VALUES ($1, $2, $3) ON CONFLICT (uuid) DO UPDATE SET amount = EXCLUDED.amount"
	_, err := tx.ExecContext(ctx, s, wal.UUID.String(), wal.UserID, wal.Amount.Amount())
	if err != nil {
		return err
	}
	return nil
}

// TODO: refactoring
func (r *Repo) SaveWallets(ctx context.Context, tx *sqlx.Tx, wal model.Wallet, wal2 model.Wallet) error {
	s := "INSERT INTO wallets (uuid, user_id, amount) VALUES ($1, $2, $3),($4,$5,$6) ON CONFLICT (uuid) DO UPDATE SET amount = EXCLUDED.amount"
	_, err := tx.ExecContext(ctx, s, wal.UUID.String(), wal.UserID, wal.Amount.Amount(), wal2.UUID.String(), wal2.UserID, wal2.Amount.Amount())
	if err != nil {
		return err
	}
	return nil
}
