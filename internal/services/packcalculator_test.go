package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePacks(t *testing.T) {
	calculator := NewPackCalculator([]int{250, 500, 1000, 2000, 5000})

	testCases := []struct {
		order    int
		expected []Pack
	}{
		{1, []Pack{{Size: 250}}},
		{250, []Pack{{Size: 250}}},
		{251, []Pack{{Size: 500}}},
		{501, []Pack{{Size: 500}, {Size: 250}}},
		{12001, []Pack{{Size: 5000}, {Size: 5000}, {Size: 2000}, {Size: 250}}},
	}

	for _, tc := range testCases {
		packs := calculator.CalculatePacks(tc.order)
		assert.Equal(t, tc.expected, packs)
	}
}
