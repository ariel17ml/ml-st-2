package planets

import (
	"github.com/golang/geo/s1"
	// "math"
)

type Planet struct {
	name      string
	speed     s1.Angle // in grades/day
	distance  int      // in km
	clockwise bool
}

const (
	CIRCUNFERENCE_DEGREES = 360
)

func (p Planet) DaysInYear() int {
	return int(CIRCUNFERENCE_DEGREES / float64(p.speed))
}
