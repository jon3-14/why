package dns

import (
	"net"
	"time"
)

type Result struct {
	Success  bool
	IPs      []string
	Duration time.Duration
	Error    string
}

func Check(host string) Result {
	start := time.Now()

	ips, err := net.LookupHost(host)

	duration := time.Since(start)

	if err != nil {
		return Result{
			Success:  false,
			Error:    err.Error(),
			Duration: duration,
		}
	}

	return Result{
		Success:  true,
		IPs:      ips,
		Duration: duration,
	}
}
