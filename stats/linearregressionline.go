package stats

import "math"

func LinearStats(x, y []float64) (float64, float64, float64) {
	sumX := getSum(x)
	sumY := getSum(y)
	xySlice := multSlice(x, y)
	xSqSlice := multSlice(x, x)
	ySqSlice := multSlice(y, y)
	xSqSum := getSum(xSqSlice)
	ySqSum := getSum(ySqSlice)
	sumXY := getSum(xySlice)
	n := float64(len(x))

	b, a := getB1B0(n, sumX, sumY, sumXY, xSqSum)
	r := getCorr(n, sumX, sumY, sumXY, xSqSum, ySqSum)

	return b, a, r
}

// getCorr() calculate the Pearson Correlation
func getCorr(n, sumX, sumY, sumXY, xSqSum, ySqSum float64) float64 {
	nume := n*sumXY - sumX*sumY
	cal := (n*xSqSum - (sumX * sumX)) * (n*ySqSum - (sumY * sumY))
	deno := math.Sqrt(cal)

	r := nume / deno

	return r
}

func getB1B0(n, sumX, sumY, sumXY, xSqSum float64) (float64, float64) {
	numerator := (n * sumXY) - (sumX * sumY)
	denominator := (n * xSqSum) - (sumX * sumX)

	b := numerator / denominator

	a := (sumY - (b * sumX)) / n

	return b, a
}

func getSum(s []float64) float64 {
	var sum float64
	for _, n := range s {
		sum += n
	}
	return sum
}

func multSlice(x, y []float64) []float64 {
	xySlice := []float64{}
	for i := range x {
		xy := x[i] * y[i]
		xySlice = append(xySlice, xy)
	}
	return xySlice
}
