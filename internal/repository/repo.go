package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Rhymond/go-money"
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

func (Repo) CreateWallet(ctx context.Context, tx sqlx.ExecerContext, userID int64, currency string) (model.Wallet, error) {
	UUID := uuid.New()
	sql := "INSERT INTO wallets (user_id, uuid, currency) VALUES ($1, $2, $3)"
	_, err := tx.ExecContext(ctx, sql, userID, UUID.String(), currency)
	if err != nil {
		return model.Wallet{}, fmt.Errorf("cant create wallet: %w", err)
	}

	return model.Wallet{
		UUID:   UUID,
		UserID: userID,
		Amount: money.New(0, currency),
	}, nil
}

func (Repo) Event(ctx context.Context, tx sqlx.ExecerContext, userID int64, amount *money.Money, targetWallet uuid.UUID, eventType string, fromWallet *uuid.UUID) error {
	args := make([]interface{}, 0, 3)
	args = append(args, userID, amount.Amount(), amount.Currency().Code, targetWallet.String(), eventType)
	var sql string
	if fromWallet != nil {
		args = append(args, fromWallet.String())
		sql = "INSERT INTO events (user_id, amount, currency, target_wallet_uuid, type, from_wallet_uuid) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	} else {
		sql = "INSERT INTO events (user_id, amount, currency,target_wallet_uuid, type) VALUES ($1, $2, $3, $4, $5) RETURNING id"
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

func (r *Repo) GetWalletInTransaction(ctx context.Context, tx *sqlx.Tx, uuid uuid.UUID) (model.Wallet, error) {
	var val wallet
	s := "SELECT uuid, user_id, amount, currency FROM wallets WHERE uuid=$1 FOR UPDATE SKIP LOCKED"
	if err := tx.GetContext(ctx, &val, s, uuid); err != nil {
		return model.Wallet{}, err
	}

	return model.Wallet{
		UUID:   val.Uuid,
		UserID: val.UserId,
		Amount: money.New(val.Amount, val.Currency),
	}, nil
}

func (r *Repo) GetWalletsInTransaction(ctx context.Context, tx *sqlx.Tx, fromUUID, toUUID uuid.UUID) (model.Wallet, model.Wallet, error) {
	var val []wallet
	s := "SELECT uuid, user_id, amount, currency FROM wallets WHERE uuid=$1 OR uuid=$2 FOR UPDATE SKIP LOCKED"
	if err := tx.SelectContext(ctx, &val, s, fromUUID, toUUID); err != nil {
		return model.Wallet{}, model.Wallet{}, err
	}
	if len(val) != 2 {
		return model.Wallet{}, model.Wallet{}, errors.New("cant get both wallets")
	}
	var from, to wallet
	if fromUUID == val[0].Uuid {
		from = val[0]
		to = val[1]
	} else {
		to = val[0]
		from = val[1]
	}

	return model.Wallet{
			UUID:   from.Uuid,
			UserID: from.UserId,
			Amount: money.New(from.Amount, from.Currency),
		}, model.Wallet{
			UUID:   to.Uuid,
			UserID: to.UserId,
			Amount: money.New(to.Amount, to.Currency),
		}, nil
}

func (r *Repo) SaveWallet(ctx context.Context, tx *sqlx.Tx, wal model.Wallet) error {
	s := "INSERT INTO wallets (uuid, user_id, amount, currency) VALUES ($1, $2, $3, $4) ON CONFLICT (uuid) DO UPDATE SET amount = EXCLUDED.amount"
	_, err := tx.ExecContext(ctx, s, wal.UUID.String(), wal.UserID, wal.Amount.Amount(), wal.Amount.Currency().Code)
	if err != nil {
		return err
	}
	return nil
}

// TODO: refactoring
func (r *Repo) SaveWallets(ctx context.Context, tx *sqlx.Tx, wal model.Wallet, wal2 model.Wallet) error {
	s := "INSERT INTO wallets (uuid, user_id, amount, currency) VALUES ($1, $2, $3, $4),($5, $6, $7, $8) ON CONFLICT (uuid) DO UPDATE SET amount = EXCLUDED.amount"
	_, err := tx.ExecContext(ctx, s,
		wal.UUID.String(), wal.UserID, wal.Amount.Amount(), wal.Amount.Currency().Code,
		wal2.UUID.String(), wal2.UserID, wal2.Amount.Amount(), wal2.Amount.Currency().Code,
	)
	if err != nil {
		return err
	}
	return nil
}

type event struct {
	ID               int64     `db:"id"`
	TargetWalletUUID string    `db:"target_wallet_uuid"`
	WalletUUID       *string   `db:"from_wallet_uuid"`
	Amount           int64     `db:"amount"`
	Date             time.Time `db:"date"`
	Type             string    `db:"type"`
}

func (r *Repo) FindEvents(ctx context.Context, userID int64, t *string, date *time.Time) ([]model.Event, error) {
	agrs := []interface{}{userID}
	var sqlBuilder strings.Builder
	sqlBuilder.WriteString("SELECT id, from_wallet_uuid, target_wallet_uuid, amount, type, date FROM events WHERE user_id=?")
	if t != nil {
		sqlBuilder.WriteString(" AND type=?")
		agrs = append(agrs, *t)
	} else {
		sqlBuilder.WriteString(" AND type IN (?,?)")
		agrs = append(agrs, model.TransferType, model.DepositType)
	}
	if date != nil {
		sqlBuilder.WriteString(" AND date>=? AND date<?")
		agrs = append(agrs, date.Format("2006-01-02"), date.Add(time.Hour*24).Format("2006-01-02"))
	}
	rows, err := r.db.QueryxContext(ctx, r.db.Rebind(sqlBuilder.String()), agrs...)
	if err != nil {
		return nil, err
	}
	var result []model.Event
	for rows.Next() {
		var dest event
		err = rows.StructScan(&dest)
		if err != nil {
			return nil, err
		}
		TargetWalletUUID, err := uuid.Parse(dest.TargetWalletUUID)
		if err != nil {
			return nil, err
		}
		var FromWalletUUID uuid.UUID
		if dest.WalletUUID != nil {
			FromWalletUUID, err = uuid.Parse(*dest.WalletUUID)
			if err != nil {
				return nil, err
			}
		}

		result = append(result, model.Event{
			ID:               0,
			UserID:           userID,
			Amount:           money.New(dest.Amount, model.DefaultCurrency),
			TargetWalletUUID: TargetWalletUUID,
			FromWalletUUID:   FromWalletUUID,
			Type:             dest.Type,
			Date:             dest.Date,
		})
	}

	return result, nil
}
