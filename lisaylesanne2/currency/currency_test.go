package currency_test

import (
	"testing"

	"github.com/keijoraamat/IDU1550/lisaylesanne2/currency"
)

func TestCurrency_ConvertTo_Should_Return_Error_When_Currencies_Are_The_Same(t *testing.T) {
	c := currency.NewCurrency("USD", "US Dollar")
	amount := 100.0

	_, err := c.ConvertTo(amount, c)

	if err == nil {
		t.Errorf("ConvertTo() should return an error when currencies are the same")
	}
}

func TestCurrency_ConvertTo_Should_Return_Error_When_No_Exchange_Rate_Found(t *testing.T) {
	c1 := currency.NewCurrency("USD", "US Dollar")
	c2 := currency.NewCurrency("EUR", "Euro")
	amount := 100.0

	_, err := c1.ConvertTo(amount, c2)

	if err == nil {
		t.Errorf("ConvertTo() should return an error when no exchange rate is found")
	}
}

func TestCurrency_ConvertTo_Should_Return_Correct_Amount_When_Exchange_Rate_Found(t *testing.T) {
	c1 := currency.NewCurrency("USD", "US Dollar")
	c2 := currency.NewCurrency("EUR", "Euro")
	c1.SetRateTo(c2, 0.85)
	amount := 100.0

	want := 117.64705882352942
	got, err := c1.ConvertTo(amount, c2)

	if err != nil {
		t.Errorf("ConvertTo() should not return an error when exchange rate is found")
	}

	if got != want {
		t.Errorf("ConvertTo() = %f, want %f", got, want)
	}
}
