package usecase

import (
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/Berchon/stress-test-cli/internal/business/entity"
	httpService "github.com/Berchon/stress-test-cli/internal/business/service"
	"github.com/Berchon/stress-test-cli/internal/infra/report"
	"github.com/Berchon/stress-test-cli/internal/infra/statistics"
)

type loadTestUseCase struct {
	httpService       httpService.HTTPService
	statisticsService statistics.StatisticsService
	reportService     report.ReportService
}

type LoadTestUseCase interface {
	StartLoadTest(targetURL string, totalRequests int, concurrency int) error
}

func NewLoadTestUseCase(
	httpService httpService.HTTPService,
	statisticsService statistics.StatisticsService,
	reportService report.ReportService,
) LoadTestUseCase {
	return &loadTestUseCase{
		httpService:       httpService,
		statisticsService: statisticsService,
		reportService:     reportService,
	}
}

func (uc *loadTestUseCase) StartLoadTest(targetURL string, totalRequests int, concurrency int) error {
	if err := validateInput(targetURL, totalRequests, concurrency); err != nil {
		return err
	}

	startTime := time.Now()

	resultsChan := make(chan entity.RequestResult, totalRequests)
	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go uc.worker(i, targetURL, totalRequests/concurrency, resultsChan, &wg)
	}

	wg.Wait()
	close(resultsChan)

	var allResults []entity.RequestResult
	for r := range resultsChan {
		allResults = append(allResults, r)
	}

	report := uc.statisticsService.AggregateResults(allResults, time.Since(startTime), targetURL)

	uc.reportService.Print(report)
	return nil
}

func (uc *loadTestUseCase) worker(workerID int, targetURL string, requestsPerWorker int, results chan<- entity.RequestResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < requestsPerWorker; i++ {
		status, latency, err := uc.httpService.DoRequest(targetURL)
		results <- entity.RequestResult{
			StatusCode: status,
			Latency:    latency,
			Error:      err,
		}
	}
}

func validateInput(targetURL string, totalRequests, concurrency int) error {
	if targetURL == "" {
		return errors.New("URL cannot be empty")
	}

	_, err := url.ParseRequestURI(targetURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %v", err)
	}

	if totalRequests <= 0 {
		return errors.New("totalRequests must be greater than 0")
	}

	if concurrency <= 0 {
		return errors.New("concurrency must be greater than 0")
	}

	if concurrency > totalRequests {
		return errors.New("concurrency cannot be greater than totalRequests")
	}

	return nil
}
