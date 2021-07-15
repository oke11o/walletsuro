package handler

import (
	"github.com/go-openapi/runtime/middleware"

	"gitlab.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
)

func (s *Server) Deposit(params wallet.DepositParams) middleware.Responder {
	//params.UserID
	return middleware.NotImplemented("CreateWallet!!!")
}
