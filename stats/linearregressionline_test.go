package stats

import (
	"math"
	"testing"
)

func TestLinearStats(t *testing.T) {
	type test struct {
		Name   string
		X_Vals []float64
		Y_Vals []float64
		B1     float64
		B0     float64
		Corr   float64
	}
	tests := []test{
		{"Testing positive values", []float64{0, 1, 2, 3, 4, 5, 6}, []float64{18, 22, 25, 29, 24, 13, 61}, 3.928571, 11.714286, 0.5415106754},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			if b, a, r := LinearStats(tc.X_Vals, tc.Y_Vals); (math.Round(b*1000000))/1000000 != tc.B1 && (math.Round(a*1000000))/1000000 != tc.B0 && (math.Round(r*10000000000))/10000000000 != tc.Corr {
				t.Errorf("%v : want: b => %v a => %v r => %v got: b => %v a => %v r => %v", tc.Name, tc.B1, tc.B0, tc.Corr, b, a, r)
			}
		})
	}
}
