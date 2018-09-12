package number_utils

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	var testsTable = []struct {
		name     string
		number   float64
		expected float64
	}{
		{"positive number #1", 13, math.Sqrt(13)},
		{"positive number #2", 31, math.Sqrt(31)},
		{"zero", 0, 0},
		{"negative number #1", -11, 0},
		{"negative number #2", -1, 0},
	}
	precision := 0.00000001

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got, _ := Sqrt(test.number)
			if test.number > 0 {
				if got < test.expected-precision || got > test.expected+precision {
					t.Errorf("Sqrt(%g) got %g expected %g", test.number, got, test.expected)
				}
			} else {
				if got != test.expected {
					t.Errorf("Sqrt(%g) got %g expected %g", test.number, got, test.expected)
				}
			}
		})
	}
}

func TestErrNegativeSqrt_Error(t *testing.T) {
	testsTable := []struct {
		name   string
		number float64
	}{
		{"positive number #1", 13},
		{"positive number #2", 31},
		{"zero", 0},
		{"negative number #1", -11},
		{"negative number #2", -1},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			_, err := Sqrt(test.number)

			if test.number >= 0 {
				if err != nil {
					t.Errorf("For %g expected <nil> got %v", test.number, err)
				}
			} else {
				if err != ErrNegativeSqrt(test.number) {
					t.Errorf("For %g expected %v got %v", test.number, ErrNegativeSqrt(test.number), err)
				}
			}
		})
	}
}
