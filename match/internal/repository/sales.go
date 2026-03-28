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

type SalesRepo interface {
	Sales(ctx context.Context) ([]model.Sale, error)
}

type salesRepo struct {
}

func newSalesRepo() SalesRepo {
	return &salesRepo{}
}

func (c salesRepo) Sales(ctx context.Context) ([]model.Sale, error) {

	var in string = `sale_id, program_id,qty,invoice_date,invoice_number,dea,pharmacy_id
1, DSH12345,10,10/11/2024,123456789,AB111111,1001
2, DSH54321,3,10/11/2024,123456787,AB222000,1001
3, DSH54321,3,15/11/2024,123456787,AB222000,1001
4, DSH11111,5,25/12/2024,111111133,AB300000,2000
5, DSH12345,1,21/12/2024,222222222,AB111111,3020`

	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	sales := []model.Sale{}

	for i, r := range records {
		if i == 0 {
			continue
		}

		saleID, err := strconv.Atoi(strings.Trim(r[0], " "))
		if err != nil {
			slog.Error("invalid saleID", slog.Any("error", err))
			continue
		}

		programID := strings.Trim(r[1], " ")

		qty, err := strconv.Atoi(strings.Trim(r[2], " "))
		if err != nil {
			slog.Error("invalid qty", slog.Any("error", err))
			continue
		}

		invoiceDate, err := time.Parse("02/01/2006", strings.Trim(r[3], " "))
		if err != nil {
			slog.Error("invalid invoiceDate", slog.Any("error", err))
			continue
		}

		invoiceNumber := strings.Trim(r[4], " ")
		drugReinforcementAgency := strings.Trim(r[5], " ")
		pharmacyID := strings.Trim(r[6], " ")

		s := model.Sale{
			SaleID:                  saleID,
			Quantity:                qty,
			InvoiceDate:             invoiceDate,
			ProgramID:               programID,
			InvoiceNumber:           invoiceNumber,
			DrugReinforcementAgency: drugReinforcementAgency,
			PharmacyID:              pharmacyID,
		}

		sales = append(sales, s)

	}

	return sales, nil

}
