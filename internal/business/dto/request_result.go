package dto

type RequestResult struct {
	WorkerID   int
	StatusCode int
	Latency    int64
	Error      error
}
