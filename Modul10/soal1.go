package main

import "fmt"

const arrkelinci int = 1000

type beratKelinci [arrkelinci] float64

func terkecil(T beratKelinci, n int) float64 {
	var min float64 = T[0]
	var i int = 1

	for i < n {
		if min > T[i] {
			min = T[i]
		}
		i = i + 1
	}
	return(min)
}

func terbesar(T beratKelinci, n int) float64 {
	var max float64 = T[0]
	var i int = 1

	for i < n {
		if max < T[i] {
			max = T[i]
		}
		i = i + 1
	}
	return(max)
}

func main() {
	var n int
	var data beratKelinci

	fmt.Scan(&n)

	if n > arrkelinci {
		n = arrkelinci
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	if n > 0 {
		fmt.Println(terkecil(data, n), terbesar(data, n))
	}
}