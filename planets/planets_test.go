package planets

import (
	"testing"
)

func TestDaysInYear(t *testing.T) {
	p := Planet{"Htrae", 1.0, 500, true}
	value := p.DaysInYear()
	expected := CIRCUNFERENCE_DEGREES

	if value != expected {
		t.Errorf("Days in year for single angle rotation is incorrect. Got: %d, expected %d", value, expected)
	}
}
