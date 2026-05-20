package tcp

import (
	"fmt"
	"net"
	"time"
)

type Result struct {
	Success bool
	Duration time.Duration
	Error string
}

func Check(host string, port int) Result {
	address := fmt.Sprintf("%s:%d", host, port)

	start := time.Now()

	conn, err := net.DialTimeout("tcp", address, 5*time.Second)

	duration := time.Since(start)

	if err != nil {
		return Result{
			Success: false,
			Error: err.Error(),
			Duration: duration,
		}
	}

	conn.Close()

	return Result{
		Success: true,
		Duration: duration,
	}
}