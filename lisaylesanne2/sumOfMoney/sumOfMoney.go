package sumofmoney

import (
	"errors"
	"fmt"

	"github.com/keijoraamat/IDU1550/lisaylesanne2/currency"
)

type SumOfMoney struct {
	amount   float64
	currency currency.Currency
}

func NewSumOfMoney(amount float64, currency currency.Currency) SumOfMoney {
	return SumOfMoney{amount: amount, currency: currency}
}

func (s *SumOfMoney) ToString() string {
	return s.currency.Code() + " " + fmt.Sprintf("%.2f", s.amount)
}

func (s *SumOfMoney) Amount() float64 {
	return s.amount
}

func (s *SumOfMoney) Currency() currency.Currency {
	return s.currency
}

func (s *SumOfMoney) Increase(amount float64) error {

	if amount < 0 {
		return errors.New("Amount to increase must be positive")
	}

	s.amount += amount
	return nil
}

func (s *SumOfMoney) Decrease(amount float64) error {

	if amount < 0 {
		return errors.New("Amount to decrease must be positive")
	}

	if amount > s.amount {
		return errors.New("Amount to decrease must not be larger than amount")
	}

	s.amount -= amount
	return nil
}

func (s *SumOfMoney) ConvertTo(currency currency.Currency) (newSumOfMoney SumOfMoney, err error) {

	c := s.Currency()
	sum, conversionErr := c.ConvertTo(s.Amount(), currency)
	if conversionErr != nil {
		return SumOfMoney{}, conversionErr
	}

	return NewSumOfMoney(sum, currency), nil
}
