package service

import (
	"context"

	"github.com/oke11o/walletsuro/internal/model"
)

func New(r repo) *Service {
	return &Service{repo: r}
}

type Service struct {
	repo repo
}

func (s Service) CreateWallet(ctx context.Context, userID int64) (model.Wallet, error) {

	return s.repo.CreateWallet(ctx, userID)
}
