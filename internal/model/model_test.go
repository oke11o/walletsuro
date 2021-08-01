package model

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestWallet_Deposit(t *testing.T) {
	tests := []struct {
		name       string
		wal        Wallet
		deposit    *money.Money
		wantAmount string
		wantErr    bool
	}{
		{
			name: "invalid currency",
			wal: Wallet{
				Amount: money.New(100, DefaultCurrency),
			},
			deposit:    money.New(34, money.AED),
			wantAmount: "",
			wantErr:    true,
		},
		{
			name: "invalid argument",
			wal: Wallet{
				Amount: money.New(100, DefaultCurrency),
			},
			deposit:    nil,
			wantAmount: "",
			wantErr:    true,
		},
		{
			name: "invalid receiver",
			wal: Wallet{
				Amount: nil,
			},
			deposit:    money.New(100, DefaultCurrency),
			wantAmount: "",
			wantErr:    true,
		},
		{
			name: "",
			wal: Wallet{
				Amount: money.New(100, DefaultCurrency),
			},
			deposit:    money.New(1234, DefaultCurrency),
			wantAmount: "$13.34",
			wantErr:    false,
		},
		{
			name: "",
			wal: Wallet{
				UUID:   uuid.UUID{},
				UserID: 0,
				Amount: money.New(0, DefaultCurrency),
			},
			deposit:    money.New(34, DefaultCurrency),
			wantAmount: "$0.34",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := tt.wal
			err := w.Deposit(tt.deposit)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantAmount, w.Amount.Display())
			}
		})
	}
}

func TestWallet_Withdraw(t *testing.T) {
	tests := []struct {
		name       string
		wal        Wallet
		withdraw   *money.Money
		wantAmount string
		wantErr    bool
	}{
		{
			name: "invalid currency",
			wal: Wallet{
				Amount: money.New(100, DefaultCurrency),
			},
			withdraw:   money.New(34, money.AED),
			wantAmount: "",
			wantErr:    true,
		},
		{
			name: "invalid argument",
			wal: Wallet{
				Amount: money.New(100, DefaultCurrency),
			},
			withdraw:   nil,
			wantAmount: "",
			wantErr:    true,
		},
		{
			name: "invalid receiver",
			wal: Wallet{
				Amount: nil,
			},
			withdraw:   money.New(100, DefaultCurrency),
			wantAmount: "",
			wantErr:    true,
		},
		{
			name: "",
			wal: Wallet{
				Amount: money.New(100, DefaultCurrency),
			},
			withdraw:   money.New(34, DefaultCurrency),
			wantAmount: "$0.66",
			wantErr:    false,
		},
		{
			name: "",
			wal: Wallet{
				Amount: money.New(0, DefaultCurrency),
			},
			withdraw:   money.New(34, DefaultCurrency),
			wantAmount: "-$0.34",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := tt.wal
			err := w.Withdraw(tt.withdraw)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantAmount, w.Amount.Display())
			}
		})
	}
}

func TestWallet_IsEnough(t *testing.T) {
	tests := []struct {
		name   string
		wal    Wallet
		amount *money.Money
		want   bool
	}{
		{
			name: "invalid currency",
			wal: Wallet{
				Amount: money.New(100, money.AED),
			},
			amount: money.New(34, DefaultCurrency),
			want:   false,
		},
		{
			name: "invalid receiver",
			wal: Wallet{
				Amount: nil,
			},
			amount: money.New(34, DefaultCurrency),
			want:   false,
		},
		{
			name: "invalid argument",
			wal: Wallet{
				Amount: money.New(100, DefaultCurrency),
			},
			amount: nil,
			want:   false,
		},
		{
			name: "not enough",
			wal: Wallet{
				Amount: money.New(100, DefaultCurrency),
			},
			amount: money.New(134, DefaultCurrency),
			want:   false,
		},
		{
			name: "success",
			wal: Wallet{
				Amount: money.New(1000, DefaultCurrency),
			},
			amount: money.New(134, DefaultCurrency),
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := tt.wal
			isEnough := w.IsEnough(tt.amount)
			require.Equal(t, tt.want, isEnough)
		})
	}
}
