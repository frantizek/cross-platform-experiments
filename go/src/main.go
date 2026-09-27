package main

import (
	"fmt"
	"runtime"
	"time"
)

func printSystemInfo() {
	fmt.Println("--- System Info ---")
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Current Time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("-------------------")
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Hello from Go!")

	printSystemInfo()

	fmt.Println("\nCalculating Fibonacci sequence...")
	for i := 1; i <= 10; i++ {
		result := fibonacci(i)
		fmt.Printf("Fib(%d) = %d\n", i, result)
	}

	fmt.Println("\nExiting...")
}
