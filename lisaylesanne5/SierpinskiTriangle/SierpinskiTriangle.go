package sierpinskitriangle

import (
	"math"

	Point "github.com/keijoraamat/IDU1550/lisaylesanne5/Point"
)

type SierpinskiTriangle struct {
	Points []Point.Point
}

func (st *SierpinskiTriangle) GeneratePoints(centre Point.Point, sideLength int) []Point.Point {
	if sideLength <= 1 {
		st.Points = append(st.Points, centre)
		return st.Points
	}

	halfSide := sideLength / 2

	st.Points = append(st.Points, st.GeneratePoints(centre, halfSide)...)
	st.Points = append(st.Points, st.GeneratePoints(Point.Point{centre.X + halfSide, centre.Y}, halfSide)...)
	st.Points = append(st.Points, st.GeneratePoints(Point.Point{centre.X + halfSide/2, centre.Y - int(math.Sqrt(3))*halfSide/2}, halfSide)...)

	return st.Points
}
