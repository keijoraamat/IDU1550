package PL

import (
	"net/http"
)

func DeleteItemHandler(service ItemServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		service.DeleteItem(id)
		w.WriteHeader(http.StatusOK)
	}
}
