package main

import "fmt"

const arrikan int = 1000

type beratikan [arrikan]float64

func main() {
	var x, y int
	var berat beratikan

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	var totalsemuaberat float64
	var jumlahwadah int

	for i := 0; i < x; i += y {
		var beratwadah float64
		for j := 0; j < y && i+j < x; j++ {
			beratwadah += berat[i+j]
		}
		fmt.Printf("%g ", beratwadah)
		totalsemuaberat += beratwadah
		jumlahwadah++
	}
	fmt.Println()

	if jumlahwadah > 0 {
		rataRata := totalsemuaberat / float64(jumlahwadah)
		fmt.Printf("Rata-ratanya %g\n", rataRata)
	}

}
