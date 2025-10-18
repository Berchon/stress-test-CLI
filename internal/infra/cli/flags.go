package cli

import "fmt"

func ValidateParams(url string, requests int, concurrency int) error {
	if url == "" {
		return fmt.Errorf("URL cannot be empty")
	}
	if requests <= 0 {
		return fmt.Errorf("requests must be a positive integer")
	}
	if concurrency <= 0 {
		return fmt.Errorf("concurrency must be a positive integer")
	}
	return nil
}
