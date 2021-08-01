package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
)

func (s *Server) Info(_ wallet.InfoParams) middleware.Responder {
	return middleware.NotImplemented("Fuck!!!")
}
