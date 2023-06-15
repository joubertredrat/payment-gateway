package infra_test

import (
	"joubertredrat/transaction-ms/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewListResponse(t *testing.T) {
	pageExpected := uint(5)
	itemsPerPageExpected := uint(50)
	dataExpected := []infra.AmountResponse{
		{
			Cents: 1250,
			BRL:   "R$12,50",
		},
	}

	listResponse := infra.NewListResponse(pageExpected, itemsPerPageExpected, dataExpected)
	assert.Equal(t, pageExpected, listResponse.Metadata.Pagination.Page)
	assert.Equal(t, itemsPerPageExpected, listResponse.Metadata.Pagination.ItemsPerPage)
	assert.Equal(t, dataExpected, listResponse.Data)
}
