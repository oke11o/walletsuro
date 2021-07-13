package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"gitlab.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
)

func (s *Server) Info(params wallet.InfoParams) middleware.Responder {
	return middleware.NotImplemented("Fuck!!!")
}
