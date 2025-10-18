package report

import (
	"fmt"
	"time"

	"github.com/Berchon/stress-test-cli/internal/business/dto"
)

type ReportService interface {
	Print(report *dto.Report)
}

type consoleReportService struct{}

func NewConsoleReportService() ReportService {
	return &consoleReportService{}
}

func (r *consoleReportService) Print(report *dto.Report) {
	if report == nil {
		fmt.Println("No report data to display.")
		return
	}

	fmt.Println()
	fmt.Println(" Stress Test Report")
	fmt.Println("--------------------")
	fmt.Printf(" Target URL:           %s\n", report.TargetURL)
	fmt.Printf(" Total Duration:       %.3fs\n", report.Duration.Seconds())
	fmt.Printf(" Total Requests:       %d\n\n", report.TotalRequests)

	fmt.Printf("Successful Requests:  %d (HTTP 200)\n\n", report.SuccessfulRequests)

	fmt.Println("Failed Requests:")
	fmt.Println("  Status Code |  Count")
	fmt.Println(" ------------------------")
	fmt.Printf("  400         | %8d\n", report.FailedRequests.HTTP400)
	fmt.Printf("  500         | %8d\n", report.FailedRequests.HTTP500)
	fmt.Printf("  Others      | %8d\n", report.FailedRequests.Others)
	fmt.Printf("  Errors      | %8d\n\n", report.FailedRequests.Errors)

	fmt.Println("Latency (ms):")
	fmt.Println("  Status Code |   Avg    |   Max    |   Min    |  StdDev")
	fmt.Println(" ---------------------------------------------------------")

	statusOrder := []string{"200", "400", "500", "Others", "Total", "Errors"}
	for _, code := range statusOrder {
		stats := report.Latency[code]
		fmt.Printf("  %-11s | %8.1f | %8.1f | %8.1f | %8.1f\n",
			code, stats.Average, float64(stats.Max), float64(stats.Min), stats.StdDev)
	}
	fmt.Println()

	fmt.Println("Percentiles (ms):")
	fmt.Println("  Status Code |   P50    |   P90    |   P95    |   P99")
	fmt.Println(" ---------------------------------------------------------")

	for _, code := range statusOrder {
		stats := report.Latency[code]
		fmt.Printf("  %-11s | %8.1f | %8.1f | %8.1f | %8.1f\n",
			code, stats.P50, stats.P90, stats.P95, stats.P99)
	}

	fmt.Println()
	fmt.Printf("Report generated at:  %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()
}
