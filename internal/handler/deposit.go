package handler

import (
	"log"

	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/runtime/middleware"
	uuid2 "github.com/google/uuid"
	"github.com/oke11o/walletsuro/internal/generated/models"
	"github.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
	"github.com/oke11o/walletsuro/internal/model"
)

func (s *Server) Deposit(params wallet.DepositParams) middleware.Responder {
	amount := model.NewMoney(params.Body.Amount, model.DefaultCurrency)
	walletUUID, err := uuid2.Parse(params.Body.WalletUUID.String())
	if err != nil {
		return wallet.NewDepositInternalServerError().WithPayload(&models.SimpleResponse{
			Message: "неверный wallet uuid",
			Status:  400,
		})
	}
	wal, err := s.service.Deposit(params.HTTPRequest.Context(), params.XUserID, walletUUID, amount)
	if err != nil {
		log.Println(err)
		return wallet.NewDepositInternalServerError().WithPayload(&models.SimpleResponse{
			Message: "ошибка сервера",
			Status:  500,
		})
	}

	return wallet.NewDepositOK().WithPayload(
		&models.Wallet{
			Amount:     wal.Amount.Amount(),
			WalletUUID: strfmt.UUID(wal.UUID.String()),
		},
	)
}
