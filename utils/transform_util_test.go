package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteToMap(t *testing.T) {

	// Test case 1: Valid JSON
	validJSON := []byte(`{"key1": "value1", "key2": 42}`)
	expected := map[string]interface{}{"key1": "value1", "key2": float64(42)}
	result, err := ByteToMap(validJSON)
	if assert.NoError(t, err, "Unexpected error for valid JSON") {
		assert.Equal(t, expected, result, "Result does not match expected value for valid JSON")
	}

	// Test case 2: Invalid JSON
	invalidJSON := []byte(`{"key1": "value1", "key2": "value2"`)
	_, err = ByteToMap(invalidJSON)
	assert.Error(t, err, "Expected an error for invalid JSON")

}
