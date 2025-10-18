package entity

// RequestResult represents the result of a single HTTP request
type RequestResult struct {
	StatusCode int
	Latency    int64
	Error      error
	// TODO: Add additional fields if needed (response size, body snippet, etc.)
}
