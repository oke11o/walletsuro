package handler

import (
	"bytes"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gocarina/gocsv"
	"github.com/oke11o/walletsuro/internal/generated/models"
	"github.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
)

func (s *Server) Report(params wallet.ReportParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	var date *time.Time
	if params.Date != nil {
		t := time.Time(*params.Date)
		date = &t
	}
	events, err := s.service.Report(ctx, params.XUserID, params.Type, date)
	if err != nil {
		return wallet.NewReportInternalServerError().WithPayload(&models.SimpleResponse{
			Message: "internal server error",
			Status:  500,
		})
	}
	var out []byte
	buffer := bytes.NewBuffer(out)
	if err := gocsv.Marshal(events, buffer); err != nil {
		return wallet.NewReportInternalServerError().WithPayload(&models.SimpleResponse{
			Message: "internal server error",
			Status:  500,
		})
	}

	return wallet.NewReportOK().WithPayload(buffer.Bytes())
}
