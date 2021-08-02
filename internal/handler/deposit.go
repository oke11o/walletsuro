package handler

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	"github.com/oke11o/walletsuro/internal/generated/models"
	"github.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
)

func (s *Server) Deposit(params wallet.DepositParams) middleware.Responder {
	walletUUID, err := uuid.Parse(params.Body.WalletUUID.String())
	if err != nil {
		return wallet.NewDepositInternalServerError().WithPayload(&models.SimpleResponse{
			Message: "неверный wallet uuid",
			Status:  400,
		})
	}
	wal, err := s.service.Deposit(params.HTTPRequest.Context(), params.XUserID, walletUUID, params.Body.Amount)
	if err != nil {
		log.Println(err)
		return wallet.NewDepositInternalServerError().WithPayload(&models.SimpleResponse{
			Message: "ошибка сервера",
			Status:  500,
		})
	}

	return wallet.NewDepositOK().WithPayload(
		&models.Wallet{
			Amount:     wal.Amount.AsMajorUnits(),
			Currency:   wal.Amount.Currency().Code,
			WalletUUID: strfmt.UUID(wal.UUID.String()),
		},
	)
}
