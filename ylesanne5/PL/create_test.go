package PL_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keijoraamat/IDU1550/ylesanne5/PL"
	"github.com/keijoraamat/IDU1550/ylesanne5/models"
)

func TestCreateItemHandler(t *testing.T) {
	// Loome uue teenuse mocki
	service := &MockService{}

	// Loome uue HTTP taotluse POST meetodiga
	item := &models.Item{Name: "item1", Price: 10}
	itemJson, _ := json.Marshal(item)
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(itemJson))
	if err != nil {
		t.Fatal(err)
	}

	// Loome uue HTTP vastuse salvestaja
	rr := httptest.NewRecorder()

	// Kutsume handlerit
	handler := PL.CreateItemHandler(service)
	handler.ServeHTTP(rr, req)

	// Kontrollime, kas saime õige HTTP staatuse koodi
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Kontrollime, kas saime õige sisu
	expectedId := 1
	var actualId int
	json.Unmarshal(rr.Body.Bytes(), &actualId)
	if actualId != expectedId {
		t.Errorf("handler returned unexpected body: got %v want %v", actualId, expectedId)
	}
}
