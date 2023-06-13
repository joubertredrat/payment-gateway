package domain_test

import (
	"joubertredrat/transaction-ms/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaginationCriteria(t *testing.T) {
	tests := []struct {
		name                       string
		page                       uint
		itemsPerPage               uint
		paginationCriteriaExpected domain.PaginationCriteria
		errExpected                error
	}{
		{
			name:         "test new pagination criteria with valid data",
			page:         7,
			itemsPerPage: 50,
			paginationCriteriaExpected: domain.PaginationCriteria{
				Page:         7,
				ItemsPerPage: 50,
			},
			errExpected: nil,
		},
		{
			name:                       "test new pagination criteria with invalid page",
			page:                       0,
			itemsPerPage:               50,
			paginationCriteriaExpected: domain.PaginationCriteria{},
			errExpected:                domain.NewErrPaginationCriteriaPage(0),
		},
		{
			name:                       "test new pagination criteria with invalid items per page",
			page:                       7,
			itemsPerPage:               190,
			paginationCriteriaExpected: domain.PaginationCriteria{},
			errExpected:                domain.NewErrPaginationCriteriaItemsPerPage(190),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			paginationCriteriaGot, errGot := domain.NewPaginationCriteria(
				test.page,
				test.itemsPerPage,
			)
			assert.Equal(t, test.paginationCriteriaExpected, paginationCriteriaGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}

func TestIsValidItemsPerPage(t *testing.T) {
	tests := []struct {
		name           string
		itemsPerPage   uint
		returnExpected bool
	}{
		{
			name:           "Test is valid items per page with valid value",
			itemsPerPage:   50,
			returnExpected: true,
		},
		{
			name:           "Test is valid items per page with less than expected value",
			itemsPerPage:   5,
			returnExpected: false,
		},
		{
			name:           "Test is valid items per page with more than expected value",
			itemsPerPage:   150,
			returnExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			returnGot := domain.IsValidItemsPerPage(test.itemsPerPage)
			assert.Equal(t, test.returnExpected, returnGot)
		})
	}
}
