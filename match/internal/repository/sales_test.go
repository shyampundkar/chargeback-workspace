package repository

import (
	"context"
	"testing"
	"time"
)

func TestSales(t *testing.T) {
	repo := newSalesRepo()
	sales, err := repo.Sales(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(sales) != 5 {
		t.Fatalf("got %d sales, want 5", len(sales))
	}

	tests := []struct {
		name          string
		index         int
		wantID        int
		wantProgramID string
		wantQty       int
		wantInvoiceNo string
		wantDate      string
		wantDEA       string
		wantPharmacy  string
	}{
		{
			name:          "first record",
			index:         0,
			wantID:        1,
			wantProgramID: "DSH12345",
			wantQty:       10,
			wantInvoiceNo: "123456789",
			wantDate:      "2024-11-10",
			wantDEA:       "AB111111",
			wantPharmacy:  "1001",
		},
		{
			name:          "second record",
			index:         1,
			wantID:        2,
			wantProgramID: "DSH54321",
			wantQty:       3,
			wantInvoiceNo: "123456787",
			wantDate:      "2024-11-10",
			wantDEA:       "AB222000",
			wantPharmacy:  "1001",
		},
		{
			name:          "third record",
			index:         2,
			wantID:        3,
			wantProgramID: "DSH54321",
			wantQty:       3,
			wantInvoiceNo: "123456787",
			wantDate:      "2024-11-15",
			wantDEA:       "AB222000",
			wantPharmacy:  "1001",
		},
		{
			name:          "fourth record — invoice number differs from matching chargeback",
			index:         3,
			wantID:        4,
			wantProgramID: "DSH11111",
			wantQty:       5,
			wantInvoiceNo: "111111133",
			wantDate:      "2024-12-25",
			wantDEA:       "AB300000",
			wantPharmacy:  "2000",
		},
		{
			name:          "fifth record",
			index:         4,
			wantID:        5,
			wantProgramID: "DSH12345",
			wantQty:       1,
			wantInvoiceNo: "222222222",
			wantDate:      "2024-12-21",
			wantDEA:       "AB111111",
			wantPharmacy:  "3020",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := sales[tc.index]

			if s.SaleID != tc.wantID {
				t.Errorf("SaleID: got %d, want %d", s.SaleID, tc.wantID)
			}
			if s.ProgramID != tc.wantProgramID {
				t.Errorf("ProgramID: got %q, want %q", s.ProgramID, tc.wantProgramID)
			}
			if s.Quantity != tc.wantQty {
				t.Errorf("Quantity: got %d, want %d", s.Quantity, tc.wantQty)
			}
			if s.InvoiceNumber != tc.wantInvoiceNo {
				t.Errorf("InvoiceNumber: got %q, want %q", s.InvoiceNumber, tc.wantInvoiceNo)
			}
			if s.InvoiceDate.Format(time.DateOnly) != tc.wantDate {
				t.Errorf("InvoiceDate: got %q, want %q", s.InvoiceDate.Format(time.DateOnly), tc.wantDate)
			}
			if s.DrugReinforcementAgency != tc.wantDEA {
				t.Errorf("DEA: got %q, want %q", s.DrugReinforcementAgency, tc.wantDEA)
			}
			if s.PharmacyID != tc.wantPharmacy {
				t.Errorf("PharmacyID: got %q, want %q", s.PharmacyID, tc.wantPharmacy)
			}
		})
	}
}
