package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	var i, count int

	count = 0
	for i = 1; i < 6; i++{
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scanln(&warna1, &warna2, &warna3, &warna4)

		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu"{
			count++
		}
	}
		if count != 0 {
			fmt.Print("BERHASIL: false")
		} else {
			fmt.Print("BERHASIL: true")
		}
}