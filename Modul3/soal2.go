package main

import "fmt"

func main() {
	var a, b, c, hasil1, hasil2, hasil3 int

	fmt.Print("Masukan Bilangan a b c: ")
	fmt.Scanln(&a, &b, &c)

	hasil1 = f(g(h(a)))
	hasil2 = g(h(f(b)))
	hasil3 = h(f(g(c)))

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}

func f(x int) int{
	return x * x
}

func g(x int) int{
	return x - 2
}

func h(x int) int{
	return x + 1
}
