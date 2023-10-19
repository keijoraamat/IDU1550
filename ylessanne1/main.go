package main

import (
	"fmt"
	"math"

	"github.com/keijoraamat/IDU1550/ylessanne1/point"
)

func main() {

	pp := point.NewPoint(4, 5)
	fmt.Println(pp)
	gg := 2.556789
	fmt.Println(math.Mod(gg, 2*math.Pi))
}
