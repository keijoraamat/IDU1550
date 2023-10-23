package currency

import (
	"errors"
)

type CurrencyInterface interface {
	AddRate(rate RateInterface)
	SetRateTo(currency CurrencyInterface, exchange float64)
}

type Currency struct {
	Code         string
	Name         string
	BaseCurrency *Currency
	Rates        []*Rate
}

func NewCurrency(code string, name string) *Currency {
	return &Currency{Code: code, Name: name}
}

func (c *Currency) AddBase(currency Currency) error {

	if currency.Code == c.Code && currency.Name == c.Name {
		return errors.New("Currency cannot be base to itself")
	}

	c.BaseCurrency = &currency
	return nil
}

func (c *Currency) AddRate(rate RateInterface) {
	c.Rates = append(c.Rates, &rate)
}

func (c *Currency) SetRateTo(currency CurrencyInterface, exchange float64) {
	r := NewRate(*c, currency)
	r.NewExchange(exchange)
	c.AddRate(r)
}
