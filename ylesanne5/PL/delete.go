package PL

import (
	"log"
	"net/http"
	"strings"
)

func DeleteItemHandler(service ItemServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the ID from the URL.
		path := r.URL.Path
		parts := strings.Split(path, "/")
		id := parts[len(parts)-1]

		log.Printf("delete, %v", id)
		service.DeleteItem(id)
		w.WriteHeader(http.StatusOK)
	}
}
