package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"gymshark-challenge/internal/services"
)

type PacksHandler struct {
	calculator *services.PackCalculator
}

func NewPacksHandler(calculator *services.PackCalculator) *PacksHandler {
	return &PacksHandler{calculator: calculator}
}

func (h *PacksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	items, err := strconv.Atoi(query.Get("items"))
	if err != nil || items <= 0 {
		log.Printf("Invalid number of items: %v", query.Get("items"))
		http.Error(w, "Invalid number of items", http.StatusBadRequest)
		return
	}

	packs := h.calculator.CalculatePacks(items)
	log.Printf("Items: %d, Packs: %v", items, packs)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packs)
}
