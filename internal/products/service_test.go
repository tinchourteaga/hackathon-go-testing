package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySeller(t *testing.T) {
	var expectedResult []Product
	expectedResult = append(expectedResult, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})

	repository := NewRepository()
	service := NewService(repository)

	result, err := service.GetAllBySeller("")

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}
