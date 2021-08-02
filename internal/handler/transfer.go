package handler

import (
	"errors"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	uuid2 "github.com/google/uuid"

	"github.com/oke11o/walletsuro/internal/generated/models"
	"github.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
	"github.com/oke11o/walletsuro/internal/model"
)

func (s *Server) Transfer(params wallet.TransferParams) middleware.Responder {
	fromWalletUUID, err := uuid2.Parse(params.Body.FromWalletUUID.String())
	if err != nil {
		return wallet.NewTransferBadRequest().WithPayload(&models.SimpleResponse{
			Message: "неверный wallet uuid",
			Status:  400,
		})
	}
	toWalletUUID, err := uuid2.Parse(params.Body.ToWalletUUID.String())
	if err != nil {
		return wallet.NewTransferBadRequest().WithPayload(&models.SimpleResponse{
			Message: "неверный wallet uuid",
			Status:  400,
		})
	}

	wal, err := s.service.Transfer(params.HTTPRequest.Context(), params.XUserID, fromWalletUUID, toWalletUUID, params.Body.Amount)

	if err != nil {
		log.Println(err)
		if errors.Is(err, model.ErrPermissionDeniedWallet) {
			return wallet.NewTransferForbidden().WithPayload(&models.SimpleResponse{
				Message: "Не достаточно прав",
				Status:  403,
			})
		}
		return wallet.NewTransferInternalServerError().WithPayload(&models.SimpleResponse{
			Message: "ошибка сервера",
			Status:  500,
		})
	}

	return wallet.NewTransferOK().WithPayload(
		&models.Wallet{
			Amount:     wal.Amount.AsMajorUnits(),
			Currency:   wal.Amount.Currency().Code,
			WalletUUID: strfmt.UUID(wal.UUID.String()),
		},
	)
}
