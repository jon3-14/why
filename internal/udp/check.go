package udp

import (
	"fmt"
	"net"
	"time"
)

type Result struct {
	Success  bool
	Duration time.Duration
	Error    string
	Response bool
}

func Check(host string, port int) Result {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	start := time.Now()

	conn, err := net.DialTimeout("udp", address, 5*time.Second)

	duration := time.Since(start)

	if err != nil {
		return Result{
			Success:  false,
			Error:    err.Error(),
			Duration: duration,
		}
	}

	defer func() {
		_ = conn.Close()
	}()

	_, err = conn.Write([]byte("ping"))

	if err != nil {
		return Result{
			Success:  false,
			Error:    err.Error(),
			Duration: duration,
		}
	}

	buffer := make([]byte, 1024)

	_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	_, err = conn.Read(buffer)

	response := err == nil

	return Result{
		Success:  true,
		Response: response,
		Duration: duration,
	}
}
