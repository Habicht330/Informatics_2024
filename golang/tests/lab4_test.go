package internal_test

import (
	"math"
	"testing"

	"isuct.ru/informatics2022/lab4"
)

func Test(t *testing.T) {
	tests := []struct {
		x, b, out float64
	}{
		{1.2, 2.5, 2.1819915488907005},
		{1.7, 2.5, 2.431423409645372},
		{2.2, 2.5, 2.6514509209692645},
		{2.7, 2.5, 2.8512957682165347},
		{3.2, 2.5, 3.0361244874322293},
		{3.7, 2.5, 3.2091876677695153},
	}
	for _, test := range tests {
		got := lab4.Calculator(test.x, test.b)
		if math.IsNaN(test.out) {
			if !math.IsNaN(got) {
				t.Errorf("F(%f, %f, %f) = %f, want NaN", test.x, test.b, got)
			}
		} else {
			if got != test.out {
				t.Errorf("F(%f, %f, %f) = %f, want %f", test.x, test.b, got, test.out)
			}
		}
	}
}
