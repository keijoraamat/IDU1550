package main

import (
	"fmt"

	sum "github.com/keijoraamat/IDU1550/lisaylesanne2/sumOfMoney"

	"github.com/keijoraamat/IDU1550/lisaylesanne2/currency"
)

func main() {

	cur := currency.NewCurrency("EUR", "Euro")
	cur.SetRateTo(*cur, 1.0)

	sum := sum.NewSumOfMoney(100, cur)
	fmt.Println(sum.ToString())

	sum.Increase(50)
	fmt.Println(sum.ToString())
}
