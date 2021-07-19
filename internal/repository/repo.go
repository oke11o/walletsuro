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

func (Repo) Event(ctx context.Context, tx sqlx.ExecerContext, userID int64, targetWallet uuid.UUID, eventType string, fromWallet *uuid.UUID) error {
	args := make([]interface{}, 0, 3)
	args = append(args, userID, targetWallet.String(), eventType)
	var sql string
	if fromWallet != nil {
		args = append(args, fromWallet.String())
		sql = "INSERT INTO events (user_id, target_wallet_uuid, type, from_wallet_uuid) VALUES ($1, $2, $3, $4) RETURNING id"
	} else {
		sql = "INSERT INTO events (user_id, target_wallet_uuid, type) VALUES ($1, $2, $3) RETURNING id"
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
