package service

import (
	"context"
	"log/slog"
	"match/internal/model"
	"match/internal/repository"
)

type Loader interface {
	ChargeBacks(ctx context.Context) ([]model.ChargeBack, error)
	Sales(ctx context.Context) ([]model.Sale, error)
}

type loader struct {
	Type string
}

func NewLoader(t string) Loader {
	return &loader{Type: t}
}

func (l loader) ChargeBacks(ctx context.Context) ([]model.ChargeBack, error) {
	ChargebackRepo, err := repository.CreateChargebackRepository(l.Type)

	if err != nil {
		slog.Error("Unable to create charge back repo")

	}

	return ChargebackRepo.ChargeBacks(ctx)

}

func (l loader) Sales(ctx context.Context) ([]model.Sale, error) {
	salesRepo, err := repository.CreateSaleRepository(l.Type)

	if err != nil {
		slog.Error("Unable to create charge back repo")

	}

	return salesRepo.Sales(ctx)
}

var _ Loader = loader{}
