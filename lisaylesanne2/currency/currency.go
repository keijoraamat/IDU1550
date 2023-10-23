package currency

type CurrencyInterface interface {
	AddRate(rate RateInterface)
	SetRateTo(currency CurrencyInterface, exchange float64)
}

type Currency struct {
	Code  string
	Name  string
	Rates []*Rate
}

func NewCurrency(code string, name string) *Currency {
	return &Currency{Code: code, Name: name}
}

func (c *Currency) AddRate(rate RateInterface) {
	if r, ok := rate.(*Rate); ok {
		c.Rates = append(c.Rates, r)
	}
}

func (c *Currency) SetRateTo(currency CurrencyInterface, exchange float64) {
	r := NewRate(c, currency)
	r.NewExchange(exchange)
	c.AddRate(r)
}
