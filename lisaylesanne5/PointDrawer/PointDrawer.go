package pointdrawer

import (
	"github.com/Pitrified/go-turtle"
	"github.com/keijoraamat/IDU1550/lisaylesanne5/Point"
	
)

type PointDrawer struct {
	Turtle *turtle.Turtle
	World  *turtle.World
	Size   int
}

func NewPointDrawer() *PointDrawer {
	pd := PointDrawer{}
	pd.Size = 2
	turtle := pd.Turtle
	pd.World = turtle.NewWorld(1000, 1000)

	return &pd
}

func (pd *PointDrawer) DrawPoints(points []Point.Point) {
	pd.Turtle.SetHeading(pd.Turtle.North)
	for _, point := range points {
		pd.DrawPoint(point)
	}
}

func (pd *PointDrawer) DrawPoint(point Point.Point) {
	pd.Turtle.PenUp()
	pd.Turtle.SetPos(point.X, point.Y)
	pd.Turtle.PenDown()

	pd.Turtle.SetSize(pd.Size + 3)
	pd.Turtle.SetColor(pd.Turtle.White)
	pd.Turtle.Forward(0)

	pd.Turtle.SetSize(pd.Size)
	pd.Turtle.SetColor(pd.Turtle.Green)
	pd.Turtle.Forward(0)
}
