package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary properties file for testing
	tempFilePath := "test_config.properties"
	tempFileContent := []byte("packSizes=300,600,1200,2400\n")
	err := os.WriteFile(tempFilePath, tempFileContent, 0644)
	assert.NoError(t, err)
	defer os.Remove(tempFilePath)

	cfg := LoadConfig(tempFilePath)
	expectedPackSizes := []int{300, 600, 1200, 2400}
	assert.Equal(t, expectedPackSizes, cfg.PackSizes)
}
