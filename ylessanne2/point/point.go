package point

import (
	"fmt"
	"math"
)

type Point struct {
	roh   float64
	theta float64
}

/*
Creates new point
RESULT.x p.X()=x
RESULT.y p.Y()=y
*/
func NewPoint(x, y float64) *Point {
	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	t := math.Atan2(y, x)
	return &Point{roh: r, theta: t}
}

/*
Returns point attributes in string form
*/
func (p *Point) String() string {
	return fmt.Sprintf("x: %f\ny: %f\nroh: %f\ntheta: %f\n", p.X(), p.Y(), p.Roh(), p.Theta())
}

/*
Returns abscissa
RESULT = roh*cos(theta)
*/
func (p *Point) X() float64 {
	return p.Roh() * math.Cos(p.Theta())
}

/*
Returns ordinate
RESULT = roh*sin(theta)
*/
func (p *Point) Y() float64 {
	return p.Roh() * math.Sin(p.Theta())
}

/*
Distance to origin (0.0, 0.0)
RESULT = sqrt(x*x, y*y)
*/
func (p *Point) Roh() float64 {
	return p.roh
}

/*
Angle to horizontal axis
RESULT = Atan2(p.y, p.x)
*/
func (p *Point) Theta() float64 {
	return p.theta
}

/*
Distance to other point
RESULT = sqrt((x*x- pnt.x*pnt.x), (y*y - pnt.y*pnt.y))
*/
func (p *Point) Distance(pnt *Point) float64 {
	v := p.VectorTo(*pnt)
	return v.Roh()
}

/*
Returns the Point representing the vector from self to other Point
RESULT.x = other.x() - x()
RESULT.y = other.y() - y()
*/
func (p *Point) VectorTo(pnt Point) Point {
	return *NewPoint(pnt.X()-p.X(), pnt.Y()-p.Y())
}

/*
Translate moves point by given amount. dx horisontally and dy vertically.
POST: x= old x +dx, y = old y +dy
*/
func (p *Point) Translate(dx float64, dy float64) {
	x := p.X() + dx
	y := p.Y() + dy

	p.roh = math.Sqrt(math.Pow(x, 2) + math.Sqrt(math.Pow(y, 2)))
	p.theta = math.Atan2(y, x)
}

/*
Scale by factor
POST: x = old x *factor && y = old y *factor
*/
func (p *Point) Scale(factor float64) {
	p.roh *= factor
}

/*
Rotate around origin (0.0, 0.0) by angle
POST: theta = old theta + angle
*/
func (p *Point) CentreRotate(angle float64) {
	p.theta = p.Theta() - angle
}

/*
Rotate around point by angle
POST: old vectorTo.theta + angle = new vectorTo.theta
*/
func (p *Point) Rotate(pnt Point, angle float64) {
	p.Translate(-pnt.X(), -pnt.Y())
	p.CentreRotate(angle)
	p.Translate(pnt.X(), pnt.Y())
}
