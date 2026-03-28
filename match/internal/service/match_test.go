package service

import (
	"match/internal/model"
	"testing"
	"time"
)

func date(s string) time.Time {
	t, _ := time.Parse(time.DateOnly, s)
	return t
}

func TestMatchKey(t *testing.T) {
	tests := []struct {
		name          string
		date          string
		invoiceNumber string
		programID     string
		want          string
	}{
		{
			name:          "standard input",
			date:          "2024-11-10",
			invoiceNumber: "123456789",
			programID:     "DSH12345",
			want:          "2024-11-10_123456789_DSH12345",
		},
		{
			name:          "lowercase inputs are uppercased",
			date:          "2024-11-10",
			invoiceNumber: "abc123",
			programID:     "dsh12345",
			want:          "2024-11-10_ABC123_DSH12345",
		},
		{
			name:          "mixed case inputs are uppercased",
			date:          "2024-12-25",
			invoiceNumber: "Inv001",
			programID:     "PrgXyz",
			want:          "2024-12-25_INV001_PRGXYZ",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := matchKey(tc.date, tc.invoiceNumber, tc.programID)
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestRunMatch(t *testing.T) {
	tests := []struct {
		name        string
		chargebacks []model.ChargeBack
		sales       []model.Sale
		wantCount   int
		wantCBIDs   []int
	}{
		{
			name: "single match",
			chargebacks: []model.ChargeBack{
				{ChargeBackID: 1, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-10")},
			},
			sales: []model.Sale{
				{SaleID: 1, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-10")},
			},
			wantCount: 1,
			wantCBIDs: []int{1},
		},
		{
			name: "no match — different invoice number",
			chargebacks: []model.ChargeBack{
				{ChargeBackID: 4, ProgramID: "DSH11111", InvoiceNumber: "111111111", InvoiceDate: date("2024-12-25")},
			},
			sales: []model.Sale{
				{SaleID: 4, ProgramID: "DSH11111", InvoiceNumber: "111111133", InvoiceDate: date("2024-12-25")},
			},
			wantCount: 0,
			wantCBIDs: nil,
		},
		{
			name: "no match — different program id",
			chargebacks: []model.ChargeBack{
				{ChargeBackID: 2, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-15")},
			},
			sales: []model.Sale{
				{SaleID: 3, ProgramID: "DSH54321", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-15")},
			},
			wantCount: 0,
			wantCBIDs: nil,
		},
		{
			name: "no match — different invoice date",
			chargebacks: []model.ChargeBack{
				{ChargeBackID: 6, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-19")},
			},
			sales: []model.Sale{
				{SaleID: 1, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-10")},
			},
			wantCount: 0,
			wantCBIDs: nil,
		},
		{
			name:        "no chargebacks",
			chargebacks: []model.ChargeBack{},
			sales: []model.Sale{
				{SaleID: 1, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-10")},
			},
			wantCount: 0,
			wantCBIDs: nil,
		},
		{
			name: "no sales",
			chargebacks: []model.ChargeBack{
				{ChargeBackID: 1, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-10")},
			},
			sales:     []model.Sale{},
			wantCount: 0,
			wantCBIDs: nil,
		},
		{
			name:        "both empty",
			chargebacks: []model.ChargeBack{},
			sales:       []model.Sale{},
			wantCount:   0,
			wantCBIDs:   nil,
		},
		{
			name: "partial match — only matching records returned",
			chargebacks: []model.ChargeBack{
				{ChargeBackID: 1, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-10")},
				{ChargeBackID: 2, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-15")},
				{ChargeBackID: 3, ProgramID: "DSH54321", InvoiceNumber: "123456787", InvoiceDate: date("2024-11-10")},
			},
			sales: []model.Sale{
				{SaleID: 1, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-10")},
				{SaleID: 2, ProgramID: "DSH54321", InvoiceNumber: "123456787", InvoiceDate: date("2024-11-10")},
			},
			wantCount: 2,
			wantCBIDs: []int{1, 3},
		},
		{
			name: "case insensitive — lowercase program id matches uppercase chargeback",
			chargebacks: []model.ChargeBack{
				{ChargeBackID: 1, ProgramID: "DSH12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-10")},
			},
			sales: []model.Sale{
				{SaleID: 1, ProgramID: "dsh12345", InvoiceNumber: "123456789", InvoiceDate: date("2024-11-10")},
			},
			wantCount: 1,
			wantCBIDs: []int{1},
		},
		{
			name: "matched result carries both chargeback and sale data",
			chargebacks: []model.ChargeBack{
				{ChargeBackID: 5, ProgramID: "DSH12345", InvoiceNumber: "222222222", InvoiceDate: date("2024-12-21")},
			},
			sales: []model.Sale{
				{SaleID: 5, ProgramID: "DSH12345", InvoiceNumber: "222222222", InvoiceDate: date("2024-12-21")},
			},
			wantCount: 1,
			wantCBIDs: []int{5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := runMatch(tc.chargebacks, tc.sales)

			if len(result) != tc.wantCount {
				t.Fatalf("got %d matches, want %d", len(result), tc.wantCount)
			}

			cbIDs := make(map[int]bool)
			for _, m := range result {
				if !m.IsMatch {
					t.Errorf("expected IsMatch=true for all returned records, got false for CB %d", m.ChargeBack.ChargeBackID)
				}
				cbIDs[m.ChargeBack.ChargeBackID] = true
			}

			for _, id := range tc.wantCBIDs {
				if !cbIDs[id] {
					t.Errorf("expected CB ID %d in results, not found", id)
				}
			}
		})
	}
}
