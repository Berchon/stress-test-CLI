package statistics

import (
	"time"

	"github.com/Berchon/stress-test-cli/internal/business/dto"
	"github.com/Berchon/stress-test-cli/internal/business/entity"
)

func aggregateResults(results []entity.RequestResult, duration time.Duration, targetURL string) *dto.Report {
	// 1. Classify results by status code
	grouped := groupResultsByStatus(results)

	// 2. Compute latency statistics for each status group
	latencyStats := computeLatencyStatsPerGroup(grouped)

	// 3. Count successes and failures
	successful, failed := countSuccessAndFailures(results)

	// 4. Build the final report DTO
	report := buildReport(targetURL, len(results), successful, failed, latencyStats, duration)

	return report
}

// groupResultsByStatus organizes RequestResults into map[statusCode][]latency
func groupResultsByStatus(results []entity.RequestResult) map[string][]float64 {
	grouped := make(map[string][]float64)
	for _, r := range results {
		key := classifyStatus(r)
		grouped[key] = append(grouped[key], float64(r.Latency))
	}
	return grouped
}

// classifyStatus returns a string key for a status code
func classifyStatus(r entity.RequestResult) string {
	if r.Error != nil {
		return "Errors"
	}
	switch {
	case r.StatusCode >= 200 && r.StatusCode < 300:
		return "200"
	case r.StatusCode >= 400 && r.StatusCode < 500:
		return "400"
	case r.StatusCode >= 500 && r.StatusCode < 600:
		return "500"
	default:
		return "Others"
	}
}

// computeLatencyStatsPerGroup calculates latency stats for each status group
func computeLatencyStatsPerGroup(grouped map[string][]float64) map[string]dto.LatencyStats {
	stats := make(map[string]dto.LatencyStats)
	for status, latencies := range grouped {
		stats[status] = computeStats(latencies)
	}

	var allLatencies []float64
	for _, lat := range grouped {
		allLatencies = append(allLatencies, lat...)
	}
	stats["Total"] = computeStats(allLatencies)

	return stats
}

func countSuccessAndFailures(results []entity.RequestResult) (successful int, failed dto.FailedRequestsBreakdown) {
	for _, r := range results {
		if r.Error != nil {
			failed.Errors++
			continue
		}
		switch {
		case r.StatusCode >= 200 && r.StatusCode < 300:
			successful++
		case r.StatusCode >= 400 && r.StatusCode < 500:
			failed.HTTP400++
		case r.StatusCode >= 500 && r.StatusCode < 600:
			failed.HTTP500++
		default:
			failed.Others++
		}
	}
	return
}

func buildReport(
	targetURL string,
	totalRequests int,
	successful int,
	failed dto.FailedRequestsBreakdown,
	latencyStats map[string]dto.LatencyStats,
	duration time.Duration,
) *dto.Report {
	return &dto.Report{
		TargetURL:          targetURL,
		TotalRequests:      totalRequests,
		SuccessfulRequests: successful,
		FailedRequests:     failed,
		Duration:           duration,
		Latency:            latencyStats,
	}
}
