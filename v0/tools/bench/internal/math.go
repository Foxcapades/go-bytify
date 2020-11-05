package internal

import "math"

func StdDev(vals []float64, stats *Stats) {
	var n, sum, ssq float64

	n = float64(len(vals))

	for i := range vals {
		sum += vals[i]
		ssq += vals[i] * vals[i]
	}

	stats.AvgTime = sum / n
	stats.StdDev = math.Sqrt(ssq/(n-1) - stats.AvgTime*stats.AvgTime)

	return
}