package service

import (
	"context"
	"log/slog"
	"match/internal/model"
	"strings"
	"time"
)

type Match struct {
	ChargeBack model.ChargeBack
	Sale       model.Sale
	IsMatch    bool
}

func matchKey(date, invoiceNumber, programID string) string {
	return strings.ToUpper(date) + "_" + strings.ToUpper(invoiceNumber) + "_" + strings.ToUpper(programID)
}

func runMatch(chs []model.ChargeBack, sales []model.Sale) []Match {
	index := make(map[string]Match)

	for _, c := range chs {
		key := matchKey(c.InvoiceDate.Format(time.DateOnly), c.InvoiceNumber, c.ProgramID)
		index[key] = Match{IsMatch: false, ChargeBack: c}
	}

	var matched []Match
	for _, s := range sales {
		key := matchKey(s.InvoiceDate.Format(time.DateOnly), s.InvoiceNumber, s.ProgramID)
		if m, ok := index[key]; ok {
			m.Sale = s
			m.IsMatch = true
			index[key] = m

			matched = append(matched, m)
		}
	}

	return matched
}

func PerformMatch(ctx context.Context) error {

	loader := NewLoader("csv")

	chs, err := loader.ChargeBacks(ctx)
	if err != nil {
		slog.Error("error getting chargebacks", slog.Any("error", err))
		return err
	}

	sales, err := loader.Sales(ctx)
	if err != nil {
		slog.Error("error getting sales", slog.Any("error", err))
		return err
	}

	result := runMatch(chs, sales)

	for _, k := range result {

		if k.IsMatch {
			slog.Info("match found")
			slog.Info("chargeback:", slog.Any("chargeback:", k.ChargeBack))
			slog.Info("sale:", slog.Any("sale:", k.Sale))
		}
	}

	return nil
}
