package currency

type RateInterface interface {
	Exchange() float64
	To() CurrencyInterface
	From() CurrencyInterface
	ChangeRate(from CurrencyInterface, to CurrencyInterface, newRate float64)
	SetExchange(e float64)
}

type Rate struct {
	from     CurrencyInterface
	to       CurrencyInterface
	exchange float64
}

func NewRate(from *Currency, to CurrencyInterface) *Rate {
	return &Rate{from: from, to: to}
}

func (r *Rate) To() CurrencyInterface {
	return r.to
}

func (r *Rate) From() CurrencyInterface {
	return r.from
}

func (r *Rate) Exchange() float64 {
	return r.exchange
}

func (r *Rate) SetExchange(e float64) {
	r.exchange = e
}

func (r *Rate) ChangeRate(from CurrencyInterface, to CurrencyInterface, newRate float64) {
	r.from = from
	r.to = to
	r.SetExchange(newRate)
}
