package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	pola(n)
}

func Bintang(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	Bintang(n - 1)
}

func pola(n int) {
	if n == 0 {
		return
	}

	pola(n - 1)
	Bintang(n)
	fmt.Println()
}
