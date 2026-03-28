package model

import (
	"time"
)

type ChargeBack struct {
	ChargeBackID            int       `csv:"cb_id"`
	ProgramID               string    `csv:"program_id"`
	InvoiceDate             time.Time `csv:"invoice_date"`
	InvoiceNumber           string    `csv:"invoice_number"`
	DrugReinforcementAgency string    `csv:"dea"`
	PharmacyID              string    `csv:"pharmacy_id"`
}

type Sale struct {
	SaleID                  int       `csv:"sale_id"`
	ProgramID               string    `csv:"program_id"`
	Quantity                int       `csv:"qty"`
	InvoiceDate             time.Time `csv:"invoice_date"`
	InvoiceNumber           string    `csv:"invoice_number"`
	DrugReinforcementAgency string    `csv:"dea"`
	PharmacyID              string    `csv:"pharmacy_id"`
}
