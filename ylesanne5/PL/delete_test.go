package PL_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keijoraamat/IDU1550/ylesanne5/PL"
)

func TestDeleteItemHandler(t *testing.T) {
	// Loome uue teenuse mocki
	service := &MockService{}

	// Loome uue HTTP taotluse DELETE meetodiga
	req, err := http.NewRequest("DELETE", "/items?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Loome uue HTTP vastuse salvestaja
	rr := httptest.NewRecorder()

	// Kutsume handlerit
	handler := PL.DeleteItemHandler(service)
	handler.ServeHTTP(rr, req)

	// Kontrollime, kas saime õige HTTP staatuse koodi
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
