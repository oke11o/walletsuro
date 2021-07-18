package handler

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
)

func (s *Server) Report(params wallet.ReportParams) middleware.Responder {
	return wallet.NewReportOK().WithPayload([]byte("asdfasdf"))
}
