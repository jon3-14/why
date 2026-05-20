package httpcheck

import (
	"net/http"
	"time"
)

type Result struct {
	Success    bool
	StatusCode int
	Status     string
	Duration   time.Duration
	Error      string
	Redirects  int
}

func Check(target string) Result {
	start := time.Now()

	redirects := 0

	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			redirects = len(via)

			if len(via) >= 10 {
				return http.ErrUseLastResponse
			}

			return nil
		},
	}

	resp, err := client.Get(target)

	duration := time.Since(start)

	if err != nil {
		return Result{
			Success:  false,
			Error:    err.Error(),
			Duration: duration,
		}
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	return Result{
		Success:    true,
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Duration:   duration,
		Redirects:  redirects,
	}
}
