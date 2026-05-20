package summary

import (
	"fmt"
	"time"
)

func PrintTiming(name string, duration time.Duration) {
	fmt.Printf("%-12s %v\n", name, duration)
}
