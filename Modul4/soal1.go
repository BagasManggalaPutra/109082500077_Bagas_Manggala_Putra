package main

import "fmt"

func main() {
	var a, b, c, d int
	var p1, p2, c1, c2 int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &p1)
	combination(a, c, &c1)
	fmt.Println(p1, c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)
	fmt.Println(p2, c2)
}

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n int, r int, hasil *int) {
	var fakN, fakNR int
	
	factorial(n, &fakN)
	factorial(n-r, &fakNR)   
	
	*hasil = fakN / fakNR
}
func combination(n int, r int, hasil *int) {
	var fakN, fakR, fakNR int
	
	factorial(n, &fakN)
	factorial(r, &fakR)      
	factorial(n-r, &fakNR)   
	
	*hasil = fakN / (fakR * fakNR)
}
