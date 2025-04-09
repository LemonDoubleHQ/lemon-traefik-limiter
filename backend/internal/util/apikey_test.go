package util

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGenerateRandomApiKey(t *testing.T) {
	apiKey := GenerateRandomApiKey()
	assert.Equal(t, len(apiKey), 64)
	assert.Equal(t, apiKey[:7], "lm-api-")
}
