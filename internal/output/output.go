package output

import "fmt"

func Section(title string) {
	fmt.Println()
	fmt.Println(title)
}

func Success(format string, args ...interface{}) {
	fmt.Printf("✓ "+format+"\n", args...)
}

func Failure(format string, args ...interface{}) {
	fmt.Printf("✗ "+format+"\n", args...)
}

func Warning(format string, args ...interface{}) {
	fmt.Printf("! "+format+"\n", args...)
}
