package statistics

import (
	"time"

	"github.com/Berchon/stress-test-cli/internal/business/dto"
	"github.com/Berchon/stress-test-cli/internal/business/entity"
)

type StatisticsService interface {
	AggregateResults(results []entity.RequestResult, duration time.Duration, targetURL string) *dto.Report
}

type statisticsService struct{}

func NewStatisticsService() StatisticsService {
	return &statisticsService{}
}

func (s *statisticsService) AggregateResults(results []entity.RequestResult, duration time.Duration, targetURL string) *dto.Report {
	return aggregateResults(results, duration, targetURL)
}
