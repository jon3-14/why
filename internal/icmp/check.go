package icmp

import (
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Result struct {
	Success  bool
	Duration time.Duration
	Error    string
	Output   string
}

func Check(host string) Result {
	var cmd *exec.Cmd

	start := time.Now()

	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", host)
	} else {
		cmd = exec.Command("ping", "-c", "1", host)
	}

	output, err := cmd.CombinedOutput()

	duration := time.Since(start)

	if err != nil {
		return Result{
			Success:  false,
			Error:    err.Error(),
			Duration: duration,
			Output:   string(output),
		}
	}

	return Result{
		Success:  true,
		Duration: duration,
		Output:   strings.TrimSpace(string(output)),
	}
}
