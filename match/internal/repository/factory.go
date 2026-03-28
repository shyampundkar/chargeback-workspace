package repository

import "fmt"

func CreateChargebackRepository(t string) (ChargebackRepo, error) {

	switch t {
	case "csv":
		return newChargebackRepo(), nil
	default:
		return nil, fmt.Errorf("repository not supported")
	}
}

func CreateSaleRepository(t string) (SalesRepo, error) {

	switch t {
	case "csv":
		return newSalesRepo(), nil
	default:
		return nil, fmt.Errorf("repository not supported")
	}

}
