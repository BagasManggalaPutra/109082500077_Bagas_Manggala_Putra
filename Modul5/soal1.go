package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	fmt.Printf("Deret Fibonacci suku ke-%d:\n", n)
	for i := 0; i <= n; i++ {
		fmt.Printf("Suku ke-%d = %d, ", i, fibonacci(i))
	}
	fmt.Println()
}