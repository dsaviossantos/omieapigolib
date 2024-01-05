package omie_api

import (
	"bytes"
	"log/slog"
	"net/http"
)

func RequestWithRetry(url string, body []byte, maxRetries int) *http.Response {
	b := bytes.NewReader(body)
	for i := 0; i < maxRetries; i++ {
		r, err := http.Post(url, "application/json", b)
		if err == nil {
			slog.Info("Response:", "Status", r.Status, "StatusCode", r.StatusCode)
			return r
		} else {
			if i == maxRetries {
				slog.Error("Max Retries Reached", "LOCATION", "utils.go:12", "MSG", err)
				panic(err)
			}
			i += 1
			slog.Warn("Request error: Retrying...", "LOCATION", "utils.go:12", "MSG", err)
		}
		defer r.Body.Close()
	}
	return nil
}
