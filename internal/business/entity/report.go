package entity

// Report contains aggregated results of the stress test
type Report struct {
	TotalRequests  int
	Successful     int
	Failed         int
	AverageLatency int64
	MeanDeviation  int64
	MaxLatency     int64
	MinLatency     int64
	// TODO: Add percentiles, errors by type, status code breakdown, etc.
}
