package main

import "fmt"

func main() {
	var brparsel, kg, sisagram, biayagram, biayakg, total,
		biayagramakhir int
	fmt.Print("Berat Parsel(gram): ")
	fmt.Scan(&brparsel)

	kg = brparsel / 1000
	sisagram = brparsel % 1000
	biayakg = kg * 10000

	if kg > 10 && sisagram < 500 {
		biayagram = 0
		biayagramakhir = sisagram * 15
	} else if kg > 10 && sisagram >= 500 {
		biayagram = 0
		biayagramakhir = sisagram * 5
	} else if sisagram < 500 {
		biayagram = sisagram * 15
		biayagramakhir = sisagram * 15
	} else {
		biayagram = sisagram * 5
		biayagramakhir = sisagram * 5
	}

	total = biayakg + biayagram
	
	fmt.Printf("Detail berat : %d kg + %d gr \n", kg, sisagram)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayakg, biayagramakhir)
	fmt.Printf("Total biaya : Rp. %d \n", total)
}
