package point_test

import (
	"math"
	"testing"

	"github.com/keijoraamat/IDU1550/ylessanne2/point"
)

// 1 * 10 in power of -6 or 0.000001
const tolerance = 1e-6

func TestPoint_X_Should_Return_Abscissa_Of_Point(t *testing.T) {
	p := point.NewPoint(7.778, 5.0)
	want := 7.778

	if got := p.X(); got != want {
		t.Errorf("X() = %f, want %f", got, want)
	}
}

func TestPoint_Y_Should_Return_Ordinate_Of_Point(t *testing.T) {
	p := point.NewPoint(1.1, -89.076)
	want := -89.076

	if math.Abs(p.Y()-want) > tolerance {
		t.Errorf("W() = %f, want %f", p.Y(), want)
	}
}

func TestPoint_Distance_To_Origin_Should_Be_As_Expected(t *testing.T) {
	p := point.NewPoint(0.0, 5.0)
	want := 5.0

	if got := p.Roh(); got != want {
		t.Errorf("Roh() = %f, want %f", got, want)
	}

	p2 := point.NewPoint(-5.0, 0.0)
	want2 := 5.0

	if got2 := p2.Roh(); got2 != want2 {
		t.Errorf("Roh() = %f, want %f", got2, want2)
	}
}

func TestPoint_Angle_To_Horizontal_Axis_Should_Be_As_Expected(t *testing.T) {
	p := point.NewPoint(0.0, 0.0)
	want := 0.0

	if got := p.Theta(); got != want {
		t.Errorf("Theta() = %f, want %f", got, want)
	}

	p2 := point.NewPoint(-5.0, 0.0)
	want2 := math.Pi

	if got2 := p2.Theta(); got2 != want2 {
		t.Errorf("Theta() = %f, want %f", got2, want2)
	}
}

func TestPoint_VectorTo_Should_Return_Point_Representing_Vector_To_Other_Point(t *testing.T) {
	p := point.NewPoint(1.0, 0.0)
	other := point.NewPoint(2.0, 3.0)
	want := point.NewPoint(1.0, 3.0)

	if got := p.VectorTo(*other); got.X() != want.X() && got.Y() != want.Y() {
		t.Errorf("VectorTo() = %f, want %f", got, want)
	}
}

func TestPoint_Distance_Should_Return_Distance_Between_Points_When_Given_Two_Points(t *testing.T) {
	p := point.NewPoint(0.0, 0.0)
	other := point.NewPoint(0.0, -1.0)
	want := 1.0

	if got := p.Distance(other); got != want {
		t.Errorf("Distance() = %f, want %F", got, want)
	}
}

func TestPoint_Translate_Should_Move_Point_By_Given_Step(t *testing.T) {
	p := point.NewPoint(1.0, -1.0)
	wantX := 0.5
	wantY := 0.0

	p.Translate(-0.5, 1.0)

	if math.Abs(p.X()-wantX) > tolerance || math.Abs(p.Y()-wantY) > tolerance {
		t.Errorf("Translate() = {%f, %f}, want {%f, %f}", p.X(), p.Y(), wantX, wantY)
	}
}

func TestPoint_Scale_Should_Scale_By_Given_Factor(t *testing.T) {
	p := point.NewPoint(2.0, 2.0)
	wantX := 1.0
	wantY := 1.0
	factor := 0.5

	p.Scale(factor)
	if math.Abs(p.X()-wantX) > tolerance || math.Abs(p.Y()-wantY) > tolerance {
		t.Errorf("Scale() = {%f, %f}, want {%f, %f}", p.X(), p.Y(), wantX, wantY)
	}

}

func TestPoint_CentreRotate_Should_Rotate_Around_Centre_Point_When_Given_Angle(t *testing.T) {
	p := point.NewPoint(0.0, 1.0)
	initialTheta := p.Theta()
	angle := 2 * math.Pi

	p.CentreRotate(angle)
	wantX := p.Roh() * math.Cos(initialTheta+angle)
	wantY := p.Roh() * math.Sin(initialTheta+angle)
	if p.X() != wantX || p.Y() != wantY {
		t.Errorf("CentreRotate(%f) = {%f, %f}, want {%f, %f}", angle, p.X(), p.Y(), wantX, wantY)
	}

}

func TestPoint_Rotate_Should_Rotate_Around_Given_Point_By_Angle(t *testing.T) {
	// Arrange
	p := point.NewPoint(4.0, 0.0)
	rotationP := point.NewPoint(0.0, 0.0)
	angle := math.Pi

	// Act
	p.Rotate(*rotationP, angle)

	// Assert
	initialVector := p.VectorTo(*rotationP)
	newX := initialVector.X()*math.Cos(angle) - initialVector.Y()*math.Sin(angle)
	newY := initialVector.X()*math.Sin(angle) + initialVector.Y()*math.Cos(angle)
	expectedX := rotationP.X() + newX
	expectedY := rotationP.Y() + newY

	if math.Abs(p.X()-expectedX) > tolerance || math.Abs(p.Y()-expectedY) > tolerance {
		t.Errorf("Rotate(%f) = {%f, %f}, want {%f, %f}", angle, p.X(), p.Y(), expectedX, expectedY)
	}
}
