package PL

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/keijoraamat/IDU1550/ylesanne5/BLL"
)

// Presentation Layer (JSON API)
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		// Handle the POST request
		fmt.Fprintf(w, "This is a POST request")
		var data = r.Body
		itemService := BLL.ItemService{}
		var ID = itemService.CreateItem(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ID)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}
