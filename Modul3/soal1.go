
package main

import "fmt"

func main() {
	var a, b, c, d, p1, p2, c1, c2 int
	
	fmt.Print("Masukan Bilangan a b c d: ")
	fmt.Scanln(&a, &b, &c, &d)
	
	p1 = permutation(a, c)
	c1 = combination(a, c)
	
	p2 = permutation(b, d)
	c2 = combination(b, d)
	
	if a < c || b < d {
		fmt.Println("Input tidak valid harus a >= c dan b >= d")
	} else {
		fmt.Println(p1, c1)
		fmt.Println(p2, c2)
	}
}

func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutation(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func combination(n, r int) int{
		return faktorial(n) / (faktorial(r) * faktorial(n-r))
}