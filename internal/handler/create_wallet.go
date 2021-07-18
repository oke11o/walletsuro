package handler

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
)

func (s *Server) CreateWallet(params wallet.CreateWalletParams) middleware.Responder {
	//params.UserID
	return middleware.NotImplemented("CreateWallet!!!")
}
