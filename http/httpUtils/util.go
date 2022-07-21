package httputils

import (
	"net/http"
	"time"
)

func HttpClient(duration *time.Duration) *http.Client {
	if duration == nil {
		return &http.Client{}
	}
	client := &http.Client{
		Timeout: *duration,
	}
	return client
}
