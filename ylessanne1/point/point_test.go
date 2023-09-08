package point_test

import (
	"math"
	"testing"

	"github.com/keijoraamat/IDU1550/ylessanne1/point"
)

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

func TestPoing_Angle_To_Horizontal_Axis_Should_Be_As_Expected(t *testing.T) {
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
