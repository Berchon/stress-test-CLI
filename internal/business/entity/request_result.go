package entity

type RequestResult struct {
	StatusCode int
	Latency    int64
	Error      error
}
