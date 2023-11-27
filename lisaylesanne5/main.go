package main

import (
	sierpinskitriangle "github.com/keijoraamat/IDU1550/lisaylesanne5/SierpinskiTriangle"

	fynedrawer "github.com/keijoraamat/IDU1550/lisaylesanne5/FyneDrawer"
	Point "github.com/keijoraamat/IDU1550/lisaylesanne5/Point"
)

func main() {

	centre := Point.Point{X: 0, Y: 0}
	sideLength := 8

	st := sierpinskitriangle.SierpinskiTriangle{}

	points := st.GeneratePoints(centre, sideLength)

	for i, point := range points {
		println(i, point.X, point.Y)
	}

	fd := fynedrawer.NewFyneDrawer()
	fd.DrawPoints(points)

}
