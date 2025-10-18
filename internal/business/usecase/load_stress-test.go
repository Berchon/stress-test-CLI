package usecase

import (
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/Berchon/stress-test-cli/internal/business/entity"
	httpService "github.com/Berchon/stress-test-cli/internal/business/service"
	"github.com/Berchon/stress-test-cli/internal/infra/statistics"
)

type loadTestUseCase struct {
	httpService       httpService.HTTPService
	statisticsService statistics.StatisticsService
}

type LoadTestUseCase interface {
	StartLoadTest(targetURL string, totalRequests int, concurrency int) error //(*dto.Report, error)
}

// NewLoadTestUseCase creates a new LoadTestUseCase instance
func NewLoadTestUseCase(httpService httpService.HTTPService, statisticsService statistics.StatisticsService) LoadTestUseCase {
	return &loadTestUseCase{
		httpService:       httpService,
		statisticsService: statisticsService,
	}
}

func (uc *loadTestUseCase) StartLoadTest(targetURL string, totalRequests int, concurrency int) error { //(*dto.Report, error) {
	// TODO: Validate input parameters (url format, totalRequests > 0, concurrency > 0)
	if err := validateInput(targetURL, totalRequests, concurrency); err != nil {
		return err
	}

	startTime := time.Now()

	// TODO: Create channels / worker pool to handle concurrency
	resultsChan := make(chan entity.RequestResult, totalRequests)
	var wg sync.WaitGroup

	// Launch workers
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go uc.worker(i, targetURL, totalRequests/concurrency, resultsChan, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	close(resultsChan)

	// TODO: Collect all RequestResult into a slice
	var allResults []entity.RequestResult
	for r := range resultsChan {
		allResults = append(allResults, r)
	}

	// TODO: Aggregate results into DTO Report (success/fail counts, latencies, etc.)
	_ = uc.statisticsService.AggregateResults(allResults, time.Since(startTime), targetURL)

	return nil
}

// TODO: Launch goroutines to perform HTTP requests
//
//	Each goroutine should record a RequestResult
func (uc *loadTestUseCase) worker(workerID int, targetURL string, requestsPerWorker int, results chan<- entity.RequestResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < requestsPerWorker; i++ {
		status, latency, err := uc.httpService.DoRequest(targetURL)
		results <- entity.RequestResult{
			StatusCode: status,
			Latency:    latency,
			Error:      err,
			// TODO: Optionally add response size, body snippet, etc.
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
