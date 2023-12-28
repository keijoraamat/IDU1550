package PL

import (
	"encoding/json"
	"net/http"

	"github.com/keijoraamat/IDU1550/ylesanne5/BLL"
)

// Presentation Layer (JSON API)
func GetAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	itemService := BLL.ItemService{}
	items := itemService.GetAllItems()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}
