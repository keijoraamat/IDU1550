package main

import (
	"fmt"

	sum "github.com/keijoraamat/IDU1550/lisaylesanne2/sumOfMoney"

	"github.com/keijoraamat/IDU1550/lisaylesanne2/currency"
)

func main() {

	//create currency
	cur := currency.NewCurrency("EUR", "Euro")

	sum := sum.NewSumOfMoney(100, cur)
	fmt.Println(sum.ToString())

	//add money
	sum.Increase(50)
	fmt.Println(sum.ToString())
	//subtract money
	sum.Decrease(25)
	fmt.Println(sum.ToString())

	//change exchange rate
	dollar := currency.NewCurrency("USD", "Dollar")
	rubles := currency.NewCurrency("RUB", "Rubles")
	cur.SetRateTo(dollar, 1.2)
	cur.SetRateTo(rubles, 100.05)

	sumCur := sum.Currency()
	fmt.Println(sumCur.ToString())

	dollaSum, err := sum.ConvertTo(dollar)
	if err != nil {
		fmt.Println(dollaSum.ToString())
	}
	fmt.Println(dollaSum.ToString())
}
