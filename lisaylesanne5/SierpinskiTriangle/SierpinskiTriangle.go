package sierpinskitriangle

import (
	"fmt"
	"math"

	Point "github.com/keijoraamat/IDU1550/lisaylesanne5/Point"
)

type SierpinskiTriangle struct {
	Points []Point.Point
}

func (st *SierpinskiTriangle) GeneratePoints(centre Point.Point, sideLength float64) {

	if sideLength <= 1.0 {
		return
	}
	fmt.Println("sideLength: ", sideLength)
	vertices := calculateTriangleVertices(centre, sideLength)
	st.Points = append(st.Points, vertices[:]...)

	for _, vertex := range vertices {
		st.GeneratePoints(vertex, sideLength/2)
	}

}

func calculateTriangleVertices(center Point.Point, sideLength float64) [3]Point.Point {
	R := sideLength / math.Sqrt(3)

	vertices := [3]Point.Point{
		{center.X, center.Y + R},
		{center.X - sideLength/2, center.Y - R/2},
		{center.X + sideLength/2, center.Y - R/2},
	}

	for i, vertex := range vertices {
		fmt.Printf("Vertex %d: (%f, %f)\n", i+1, vertex.X, vertex.Y)
	}
	return vertices
}
