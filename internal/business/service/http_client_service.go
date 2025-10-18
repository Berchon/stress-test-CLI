package service

import (
	"net/http"
	"time"
)

type HTTPService interface {
	DoRequest(targetURL string) (int, int64, error)
}

type httpService struct {
	client *http.Client
}

func NewHTTPService(timeout time.Duration) HTTPService {
	return &httpService{
		client: &http.Client{Timeout: timeout},
	}
}

func (s *httpService) DoRequest(targetURL string) (int, int64, error) {
	start := time.Now()

	resp, err := s.client.Get(targetURL)
	latency := time.Since(start).Milliseconds()

	if err != nil {
		return 0, latency, err
	}

	defer resp.Body.Close()
	return resp.StatusCode, latency, nil
}
