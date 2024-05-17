package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMetaPagination(t *testing.T) {

	// Test case 1: Total less than limit
	meta := CreateMetaPagination(1, 20, 15)
	assert.Equal(t, 1, meta.CurrentPage, "CurrentPage should be 1")
	assert.Equal(t, 20, meta.PerPage, "PerPage should be 20")
	assert.Equal(t, 1, meta.From, "From should be 1")
	assert.Equal(t, 15, meta.To, "To should be 15")
	assert.Equal(t, 15, meta.Total, "Total should be 15")
	assert.Equal(t, 1, meta.LastPage, "LastPage should be 1")

	// Test case 2: Total greater than limit
	meta = CreateMetaPagination(2, 20, 50)
	assert.Equal(t, 2, meta.CurrentPage, "CurrentPage should be 2")
	assert.Equal(t, 20, meta.PerPage, "PerPage should be 20")
	assert.Equal(t, 21, meta.From, "From should be 21")
	assert.Equal(t, 40, meta.To, "To should be 40")
	assert.Equal(t, 50, meta.Total, "Total should be 50")
	assert.Equal(t, 3, meta.LastPage, "LastPage should be 3")

}
