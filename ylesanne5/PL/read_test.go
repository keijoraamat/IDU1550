package PL_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keijoraamat/IDU1550/ylesanne5/PL"
	"github.com/keijoraamat/IDU1550/ylesanne5/models"
)

func TestGetAllItemsHandler(t *testing.T) {
	// Loome uue teenuse mocki
	service := &MockService{}

	// Loome uue HTTP taotluse GET meetodiga
	req, err := http.NewRequest("GET", "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Loome uue HTTP vastuse salvestaja
	rr := httptest.NewRecorder()

	// Kutsume handlerit
	handler := PL.GetAllItemsHandler(service)
	handler.ServeHTTP(rr, req)

	// Kontrollime, kas saime õige HTTP staatuse koodi
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Kontrollime, kas saime õige sisu
	expectedItems := []models.Item{{Name: "item1", Price: 10}, {Name: "item2", Price: 20}}
	var actualItems []models.Item
	json.Unmarshal(rr.Body.Bytes(), &actualItems)
	for i, item := range actualItems {
		if item != expectedItems[i] {
			t.Errorf("handler returned unexpected body: got %v want %v", item, expectedItems[i])
		}
	}
}
