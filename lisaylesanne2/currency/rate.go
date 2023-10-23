package currency

type RateInterface interface {
	ChangeRate(from CurrencyInterface, to CurrencyInterface, newRate float64)
	NewExchange(e float64)
}

type Rate struct {
	From     CurrencyInterface
	To       CurrencyInterface
	Exchange float64
}

func NewRate(from CurrencyInterface, to CurrencyInterface) *Rate {
	return &Rate{From: from, To: to}
}

func (r *Rate) ChangeRate(from CurrencyInterface, to CurrencyInterface, newRate float64) {
	r.From = from
	r.To = to
	r.Exchange = newRate
}

func (r *Rate) NewExchange(e float64) {
	r.Exchange = e
}
