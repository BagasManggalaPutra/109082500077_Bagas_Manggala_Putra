package main

import "fmt"

type Nmax[50] string

func main() {
	var klubA string
	var klubB string
	var hasil Nmax
	var n int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pertandingan := 1
	for {
		var skorA int
		var skorB int

		fmt.Print("Pertandingan ", pertandingan, " : ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
		} else {
			hasil[n] = "Draw"
		}

		n++
		pertandingan++
	}

	for i := 0; i < n; i++ {
		fmt.Println("Hasil", i+1, ":", hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}