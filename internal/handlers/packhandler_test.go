package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gymshark-challenge/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestPacksHandler(t *testing.T) {
	calculator := services.NewPackCalculator([]int{250, 500, 1000, 2000, 5000})
	handler := NewPacksHandler(calculator)

	testCases := []struct {
		items    string
		expected int
		packs    []services.Pack
	}{
		{"1", http.StatusBadRequest, nil},
		{"250", http.StatusOK, []services.Pack{{Size: 250}}},
		{"501", http.StatusOK, []services.Pack{{Size: 500}, {Size: 250}}},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest("GET", "/packs?items="+tc.items, nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, tc.expected, rr.Code)

		if rr.Code == http.StatusOK {
			var packs []services.Pack
			err := json.NewDecoder(rr.Body).Decode(&packs)
			assert.NoError(t, err)
			assert.Equal(t, tc.packs, packs)
		}
	}
}
