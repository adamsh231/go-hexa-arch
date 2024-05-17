package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidationPaginationDefault(t *testing.T) {

	// Test case 1: Valid page and limit values
	page, limit := ValidationPaginationDefault(1, 20)
	assert.Equal(t, 1, page, "Page value should be 1")
	assert.Equal(t, 20, limit, "Limit value should be 20")

	// Test case 2: Negative page value
	page, limit = ValidationPaginationDefault(-1, 20)
	assert.Equal(t, 1, page, "Page value should default to 1 for negative input")
	assert.Equal(t, 20, limit, "Limit value should remain unchanged")

	// Test case 3: Negative limit value
	page, limit = ValidationPaginationDefault(1, -20)
	assert.Equal(t, 1, page, "Page value should remain unchanged")
	assert.Equal(t, 10, limit, "Limit value should default to 10 for negative input")

}

func TestValidateStruct(t *testing.T) {

	// Test case 1: Valid struct without validation errors
	type TestData struct {
		Field1 int    `validate:"min=0"`
		Field2 string `validate:"required"`
	}
	data := TestData{Field1: 5, Field2: "test"}
	errs := ValidateStruct(data)
	assert.Empty(t, errs, "There should be no validation errors")

	// Test case 2: Struct with validation errors
	type TestData2 struct {
		Field1 int    `validate:"min=10"`
		Field2 string `validate:"required"`
	}
	data2 := TestData2{Field1: 5, Field2: ""}
	errs2 := ValidateStruct(data2)
	assert.Len(t, errs2, 2, "There should be 2 validation errors")

	// Add more test cases as needed to cover different scenarios

}
