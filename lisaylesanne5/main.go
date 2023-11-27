package main

import (
	"fmt"

	sierpinskitriangle "github.com/keijoraamat/IDU1550/lisaylesanne5/SierpinskiTriangle"

	Point "github.com/keijoraamat/IDU1550/lisaylesanne5/Point"
	pointdrawer "github.com/keijoraamat/IDU1550/lisaylesanne5/PointDrawer"
)

func main() {

	centre := Point.Point{X: 500.0, Y: 500.0}
	sideLength := 400.0

	fmt.Println("started")
	st := sierpinskitriangle.SierpinskiTriangle{}

	fmt.Println("sierpinski intialized")
	st.GeneratePoints(centre, sideLength)

	fmt.Println("points generated", len(st.Points))

	fd := pointdrawer.NewPointDrawer()
	fd.DrawPoints(st.Points)

}
