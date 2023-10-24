package sumOfMoney_test

import (
	"testing"

	"github.com/keijoraamat/IDU1550/lisaylesanne2/currency"
	sum "github.com/keijoraamat/IDU1550/lisaylesanne2/sumOfMoney"
)

func TestSumOfMoney_ConvertTo_Should_Convert_To_Given_Currency(t *testing.T) {
	// Arrange
	dollar := currency.NewCurrency("USD", "Dollar")
	euro := currency.NewCurrency("EUR", "Euro")
	dollar.SetRateTo(euro, 1.2)
	sumInDollars := sum.NewSumOfMoney(100, dollar)

	want := sum.NewSumOfMoney(120, euro)

	// Act
	got, err := sumInDollars.ConvertTo(euro)

	// Assert
	if err != nil {
		t.Errorf("ConvertTo() returned an error: %v", err)
	}
	if got.Amount() != want.Amount() && got.Currency() != want.Currency() {
		t.Errorf("ConvertTo() = %v, want %v", got, want)
	}
}
