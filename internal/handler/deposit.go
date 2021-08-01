package handler

import (
	"fmt"
	"log"

	"github.com/Rhymond/go-money"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	"github.com/oke11o/walletsuro/internal/generated/models"
	"github.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
	"github.com/oke11o/walletsuro/internal/utils"
)

func (s *Server) Deposit(params wallet.DepositParams) middleware.Responder {
	currency := money.GetCurrency(params.Body.Currency)
	if currency == nil {
		return wallet.NewDepositBadRequest().WithPayload(&models.SimpleResponse{
			Message: fmt.Sprintf("invalid currency: %s", params.Body.Currency),
			Status:  400,
		})
	}
	amount := money.New(utils.FloatWithFraction(params.Body.Amount, currency.Fraction), currency.Code)
	walletUUID, err := uuid.Parse(params.Body.WalletUUID.String())
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
			Amount:     wal.Amount.AsMajorUnits(),
			Currency:   wal.Amount.Currency().Code,
			WalletUUID: strfmt.UUID(wal.UUID.String()),
		},
	)
}
