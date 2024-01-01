package PL

import (
	"encoding/json"
	"net/http"
)

// Presentation Layer (JSON API)
func GetAllItemsHandler(service ItemServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := service.GetAllItems()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(items)
	}
}
