package main

import (
	"fmt"
	"math"
)

type arr [20]int

func semua(a arr, n int) {
	fmt.Print("Semua elemen : ")
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

func ganjil(a arr, n int) {
	fmt.Print("Index ganjil : ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(a[i], " ")
		}
	}
	fmt.Println()
}

func genap(a arr, n int) {
	fmt.Print("Index genap : ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a[i], " ")
		}
	}
	fmt.Println()
}

func kelipatanX(a arr, n int, x int) {
	fmt.Print("Index kelipatan", x, " : ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(a[i], " ")
		}
	}
	fmt.Println()
}

func hapusIndex(a *arr, n *int, indx int) {
	for i := indx; i < *n-1; i++ {
		a[i] = a[i+1]
	}
	*n = *n - 1
}

func rata2(a arr, n int) float64 {
	hasil := 0
	for i := 0; i < n; i++ {
		hasil = hasil + a[i]
	}
	rata := (float64(hasil) / float64(n))
	return rata
}

func standarDev(a arr, n int) float64 {
	rata := rata2(a, n)
	total := 0.0
	for i := 0; i < n; i++ {
		selisih := float64(a[i]) - rata
		total = total + (selisih * selisih)
	}
	hasil := math.Sqrt(total / float64(n))
	return hasil
}

func frekuensi(a arr, n int, bil int) int {
	total := 0
	for i := 0; i < n; i++ {
		if a[i] == bil {
			total++
		}
	}
	return total
}

func main() {
	var arr arr
	var n int

	fmt.Print("Masukkan jumlah elemen : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Index ke-", i, " : ")
		fmt.Scan(&arr[i])
	}
	fmt.Println()

	//a
	semua(arr, n)

	//b
	ganjil(arr, n)

	//c
	genap(arr, n)

	//d
	var x int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	kelipatanX(arr, n, x)

	//e
	var indx int
	fmt.Print("Masukkan indeks yang dihapus : ")
	fmt.Scan(&indx)
	hapusIndex(&arr, &n, indx)
	fmt.Print("Array setelah dihapus : ")
	semua(arr, n)

	//f
	fmt.Printf("Rata-rata : %.2f\n", rata2(arr, n))

	//g
	fmt.Printf("Standar deviasi : %.2f\n", standarDev(arr, n))

	//h
	var mencari int 
	fmt.Print("Cari frekuensi angka : ")
	fmt.Scan(&mencari)
	fmt.Println("Frekuensi", mencari, ":", frekuensi(arr, n, mencari), "kali")
}
