package handler

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
)

func (s *Server) CreateWallet(params wallet.CreateWalletParams) middleware.Responder {
	wal, err := s.service.CreateWallet(params.HTTPRequest.Context(), params.XUserID, params.Body.Currency)
	if err != nil {
		log.Println(err)
		return wallet.NewCreateWalletInternalServerError()
	}

	return wallet.NewCreateWalletOK().WithPayload(
		&wallet.CreateWalletOKBody{
			WalletUUID: strfmt.UUID(wal.UUID.String()),
		},
	)
}
