package dto

// RequestResultDTO is used to transfer request results to the report layer
type RequestResult struct {
	WorkerID   int // optional, for internal concurrency tracking
	StatusCode int
	Latency    int64
	Error      error
}
