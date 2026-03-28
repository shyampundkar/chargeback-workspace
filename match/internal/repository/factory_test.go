package repository

import (
	"testing"
)

func TestCreateChargebackRepository(t *testing.T) {
	tests := []struct {
		name     string
		repoType string
		wantNil  bool
		wantErr  bool
	}{
		{
			name:     "csv returns valid repo",
			repoType: "csv",
			wantNil:  false,
			wantErr:  false,
		},
		{
			name:     "unknown type returns error",
			repoType: "postgres",
			wantNil:  true,
			wantErr:  true,
		},
		{
			name:     "empty type returns error",
			repoType: "",
			wantNil:  true,
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo, err := CreateChargebackRepository(tc.repoType)

			if tc.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if tc.wantNil && repo != nil {
				t.Error("expected nil repo, got non-nil")
			}
			if !tc.wantNil && repo == nil {
				t.Error("expected non-nil repo, got nil")
			}
		})
	}
}

func TestCreateSaleRepository(t *testing.T) {
	tests := []struct {
		name     string
		repoType string
		wantNil  bool
		wantErr  bool
	}{
		{
			name:     "csv returns valid repo",
			repoType: "csv",
			wantNil:  false,
			wantErr:  false,
		},
		{
			name:     "unknown type returns error",
			repoType: "postgres",
			wantNil:  true,
			wantErr:  true,
		},
		{
			name:     "empty type returns error",
			repoType: "",
			wantNil:  true,
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo, err := CreateSaleRepository(tc.repoType)

			if tc.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if tc.wantNil && repo != nil {
				t.Error("expected nil repo, got non-nil")
			}
			if !tc.wantNil && repo == nil {
				t.Error("expected non-nil repo, got nil")
			}
		})
	}
}
