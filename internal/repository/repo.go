package repository

import (
	"context"
	"fmt"

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

func (r Repo) CreateWallet(ctx context.Context, userID int64) (model.Wallet, error) {
	UUID := uuid.New()
	sql := "INSERT INTO wallets (user_id, wallet) VALUES ($1, $2)"
	result, err := r.db.ExecContext(ctx, sql, userID, UUID.String())
	if err != nil {
		return model.Wallet{}, fmt.Errorf("cant create wallet: %w", err)
	}

	wallet := model.Wallet{
		UUID:   UUID,
		UserID: userID,
	}
	walletID, err := result.LastInsertId()
	if err != nil {
		return wallet, nil
	}
	wallet.ID = walletID
	return wallet, nil
}
