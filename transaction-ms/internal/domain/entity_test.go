package domain_test

import (
	"joubertredrat/transaction-ms/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidInstallments(t *testing.T) {
	tests := []struct {
		name           string
		installment    uint
		returnExpected bool
	}{
		{
			name:           "Test is valid installments with valid installment",
			installment:    4,
			returnExpected: true,
		},
		{
			name:           "Test is valid installments with less than minimum expected",
			installment:    0,
			returnExpected: false,
		},
		{
			name:           "Test is valid installments with more than maximum expected",
			installment:    15,
			returnExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			returnGot := domain.IsValidInstallments(test.installment)
			assert.Equal(t, test.returnExpected, returnGot)
		})
	}
}

func TestIsValidStatus(t *testing.T) {
	tests := []struct {
		name           string
		status         string
		returnExpected bool
	}{
		{
			name:           "Test is valid status with valid status",
			status:         domain.STATUS_CREATED,
			returnExpected: true,
		},
		{
			name:           "Test is valid status with invalid status",
			status:         "done",
			returnExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			returnGot := domain.IsValidStatus(test.status)
			assert.Equal(t, test.returnExpected, returnGot)
		})
	}
}
