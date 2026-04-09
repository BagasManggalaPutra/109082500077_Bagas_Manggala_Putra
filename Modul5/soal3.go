package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	mencariFaktor(n, 1)
	fmt.Println() 
}

func mencariFaktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Printf("%d ", i)
	}

	mencariFaktor(n, i+1)
}
