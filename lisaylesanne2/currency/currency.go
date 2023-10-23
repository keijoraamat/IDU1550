package currency

import (
	"errors"
	"fmt"
)

type CurrencyInterface interface {
	Code() string
	Name() string
	GetRates() []RateInterface
	AddRate(currency CurrencyInterface, rate RateInterface)
	SetRateTo(currency CurrencyInterface, exchange float64)
}

type Currency struct {
	code  string
	name  string
	Rates []RateInterface
}

func NewCurrency(code string, name string) *Currency {
	c := Currency{}
	c.SetCode(code)
	c.SetName(name)
	return &c
}

func (c *Currency) AddRate(currency CurrencyInterface, rate RateInterface) {

	c.Rates = append(c.Rates, rate)

}

func (c *Currency) SetRateTo(currency CurrencyInterface, exchange float64) {
	for _, rate := range c.GetRates() {
		to := rate.To()
		if to.Code() == currency.Code() {
			rate.SetExchange(exchange)
			return
		}
	}
	fmt.Println("Got new currency, ", currency.Code())
	newRate := NewRate(c, currency)
	newRate.SetExchange(exchange)
	c.AddRate(currency, newRate)
}

func (c *Currency) GetRates() []RateInterface {
	return c.Rates
}

func (c *Currency) ToString() string {
	return c.Code() + " " + c.Name()
}

func (c *Currency) Code() string {
	return c.code
}

func (c *Currency) Name() string {
	return c.name
}

func (c *Currency) SetCode(code string) {
	c.code = code
}

func (c *Currency) SetName(name string) {
	c.name = name
}

func (c *Currency) ConvertTo(amount float64, currency *Currency) (newAmount float64, err error) {

	if c.Code() == currency.Code() && c.Name() == currency.Name() {
		return 0, errors.New("Currency can't should not be the same")
	}

	for _, rate := range c.Rates {
		to := rate.To()
		if to.Code() == currency.Code() && to.Name() == currency.Name() {
			newAmount = amount / rate.Exchange()
			return newAmount, nil
		}
	}

	return 0, errors.New("No exchange rate found")
}
