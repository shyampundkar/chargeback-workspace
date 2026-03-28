package repository

import (
	"context"
	"encoding/csv"
	"log/slog"
	"match/internal/model"
	"strconv"
	"strings"
	"time"
)

type ChargebackRepo interface {
	ChargeBacks(ctx context.Context) ([]model.ChargeBack, error)
}

type chargeBackRepo struct {
}

func newChargebackRepo() ChargebackRepo {
	return &chargeBackRepo{}
}

func (c chargeBackRepo) ChargeBacks(ctx context.Context) ([]model.ChargeBack, error) {

	var in string = `cb_id, program_id,qty,invoice_date,invoice_number,dea,pharmacy_id
1, DSH12345,3,10/11/2024,123456789,AB111111,1001
2, DSH12345,3,15/11/2024,123456789,AB111111,1001
3, DSH54321,3,10/11/2024,123456787,AB222000,1001
4, DSH11111,5,25/12/2024,111111111,AB300000,2000
5, DSH12345,1,21/12/2024,222222222,AB111111,1001
6, DSH12345,8,19/11/2024,123456789,AB232323,1001`

	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	chs := []model.ChargeBack{}

	for i, r := range records {
		if i == 0 {
			continue
		}

		chargeBackID, err := strconv.Atoi(strings.Trim(r[0], " "))
		if err != nil {
			slog.Error("invalid chargebackId", slog.Any("error", err))
			continue
		}

		invoiceDate, err := time.Parse("02/01/2006", strings.Trim(r[3], " "))
		if err != nil {
			slog.Error("invalid invoiceDate", slog.Any("error", err))
			continue
		}

		programID := strings.Trim(r[1], " ")
		invoiceNumber := strings.Trim(r[4], " ")
		drugReinforcementAgency := strings.Trim(r[5], " ")
		pharmacyID := strings.Trim(r[6], " ")

		ch := model.ChargeBack{
			ChargeBackID:            chargeBackID,
			InvoiceDate:             invoiceDate,
			ProgramID:               programID,
			InvoiceNumber:           invoiceNumber,
			DrugReinforcementAgency: drugReinforcementAgency,
			PharmacyID:              pharmacyID,
		}

		chs = append(chs, ch)

	}

	return chs, nil

}
