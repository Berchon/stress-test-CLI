package dto

import "time"

type Report struct {
	TargetURL          string
	TotalRequests      int
	SuccessfulRequests int
	FailedRequests     FailedRequestsBreakdown

	Duration time.Duration

	Latency map[string]LatencyStats // keys: "200", "400", "500", "Others", "Total", "Errors"
}

type LatencyStats struct {
	Average float64
	Max     int64
	Min     int64
	StdDev  float64
	P50     float64
	P90     float64
	P95     float64
	P99     float64
}

type FailedRequestsBreakdown struct {
	HTTP400 int
	HTTP500 int
	Others  int
	Errors  int
}
