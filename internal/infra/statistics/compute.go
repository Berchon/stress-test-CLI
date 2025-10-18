package statistics

import (
	"sort"

	"github.com/Berchon/stress-test-cli/internal/business/dto"
	"gonum.org/v1/gonum/stat"
)

func computeStats(latencies []float64) dto.LatencyStats {
	if len(latencies) == 0 {
		return dto.LatencyStats{}
	}

	min, max, sum := computeMinMaxSum(latencies)
	average := computeAverage(sum, len(latencies))
	stdDev := computeStdDev(latencies, average)
	percentiles := computePercentiles(latencies, []int{50, 90, 95, 99})

	return dto.LatencyStats{
		Average: average,
		Max:     int64(max),
		Min:     int64(min),
		StdDev:  stdDev,
		P50:     percentiles[50],
		P90:     percentiles[90],
		P95:     percentiles[95],
		P99:     percentiles[99],
	}
}

func computeMinMaxSum(values []float64) (min, max, sum float64) {
	min = values[0]
	max = values[0]
	sum = 0
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}
	return
}

func computeAverage(sum float64, count int) float64 {
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

func computeStdDev(values []float64, mean float64) float64 {
	if len(values) == 0 {
		return 0
	}
	return stat.StdDev(values, nil)
}

func computePercentiles(values []float64, percentiles []int) map[int]float64 {
	// Values must be sorted ascending before calling stat.Quantile to avoid panic.
	sort.Float64s(values)

	result := make(map[int]float64)
	for _, p := range percentiles {
		result[p] = stat.Quantile(float64(p)/100.0, stat.Empirical, values, nil)
	}
	return result
}
