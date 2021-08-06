package handler

import (
	"bytes"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gocarina/gocsv"

	"github.com/oke11o/walletsuro/internal/generated/models"
	"github.com/oke11o/walletsuro/internal/generated/restapi/operations/wallet"
	"github.com/oke11o/walletsuro/internal/model"
)

type reportData struct {
	WalletUUID string `csv:"wallet_uuid"`
	Date       string `csv:"date"`
	Type       string `csv:"type"`
	Amount     string `csv:"amount"`
}

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

	repData := event2repData(events)

	var out []byte
	buffer := bytes.NewBuffer(out)
	if err := gocsv.Marshal(repData, buffer); err != nil {
		return wallet.NewReportInternalServerError().WithPayload(&models.SimpleResponse{
			Message: "internal server error",
			Status:  500,
		})
	}

	return wallet.NewReportOK().WithPayload(buffer.Bytes())
}

func event2repData(events []model.Event) []reportData {
	result := make([]reportData, 0, len(events))
	for _, e := range events {
		result = append(result, reportData{
			WalletUUID: e.WalletUUID.String(),
			Date:       e.Date.Format(time.RFC3339),
			Type:       e.Type,
			Amount:     e.Amount.Display(),
		})
	}
	return result
}
