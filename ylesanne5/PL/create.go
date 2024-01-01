package PL

import (
	"encoding/json"
	"net/http"

	"github.com/keijoraamat/IDU1550/ylesanne5/models"
)

func CreateItemHandler(service ItemServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item models.Item
		json.NewDecoder(r.Body).Decode(&item)
		id := service.CreateItem(&item)
		json.NewEncoder(w).Encode(id)
	}
}
