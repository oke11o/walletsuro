package handler

import (
	"github.com/go-openapi/runtime/middleware"

	"gitlab.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
)

func (s *Server) Transfer(params wallet.TransferParams) middleware.Responder {
	return middleware.NotImplemented("CreateWallet!!!")
}
