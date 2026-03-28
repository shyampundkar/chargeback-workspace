package repository

import (
	"context"
	"testing"
	"time"
)

func TestChargeBacks(t *testing.T) {
	repo := newChargebackRepo()
	chs, err := repo.ChargeBacks(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(chs) != 6 {
		t.Fatalf("got %d chargebacks, want 6", len(chs))
	}

	tests := []struct {
		name          string
		index         int
		wantID        int
		wantProgramID string
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
			wantInvoiceNo: "123456789",
			wantDate:      "2024-11-10",
			wantDEA:       "AB111111",
			wantPharmacy:  "1001",
		},
		{
			name:          "second record — same invoice number different date",
			index:         1,
			wantID:        2,
			wantProgramID: "DSH12345",
			wantInvoiceNo: "123456789",
			wantDate:      "2024-11-15",
			wantDEA:       "AB111111",
			wantPharmacy:  "1001",
		},
		{
			name:          "third record",
			index:         2,
			wantID:        3,
			wantProgramID: "DSH54321",
			wantInvoiceNo: "123456787",
			wantDate:      "2024-11-10",
			wantDEA:       "AB222000",
			wantPharmacy:  "1001",
		},
		{
			name:          "fourth record",
			index:         3,
			wantID:        4,
			wantProgramID: "DSH11111",
			wantInvoiceNo: "111111111",
			wantDate:      "2024-12-25",
			wantDEA:       "AB300000",
			wantPharmacy:  "2000",
		},
		{
			name:          "fifth record",
			index:         4,
			wantID:        5,
			wantProgramID: "DSH12345",
			wantInvoiceNo: "222222222",
			wantDate:      "2024-12-21",
			wantDEA:       "AB111111",
			wantPharmacy:  "1001",
		},
		{
			name:          "sixth record",
			index:         5,
			wantID:        6,
			wantProgramID: "DSH12345",
			wantInvoiceNo: "123456789",
			wantDate:      "2024-11-19",
			wantDEA:       "AB232323",
			wantPharmacy:  "1001",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ch := chs[tc.index]

			if ch.ChargeBackID != tc.wantID {
				t.Errorf("ChargeBackID: got %d, want %d", ch.ChargeBackID, tc.wantID)
			}
			if ch.ProgramID != tc.wantProgramID {
				t.Errorf("ProgramID: got %q, want %q", ch.ProgramID, tc.wantProgramID)
			}
			if ch.InvoiceNumber != tc.wantInvoiceNo {
				t.Errorf("InvoiceNumber: got %q, want %q", ch.InvoiceNumber, tc.wantInvoiceNo)
			}
			if ch.InvoiceDate.Format(time.DateOnly) != tc.wantDate {
				t.Errorf("InvoiceDate: got %q, want %q", ch.InvoiceDate.Format(time.DateOnly), tc.wantDate)
			}
			if ch.DrugReinforcementAgency != tc.wantDEA {
				t.Errorf("DEA: got %q, want %q", ch.DrugReinforcementAgency, tc.wantDEA)
			}
			if ch.PharmacyID != tc.wantPharmacy {
				t.Errorf("PharmacyID: got %q, want %q", ch.PharmacyID, tc.wantPharmacy)
			}
		})
	}
}
