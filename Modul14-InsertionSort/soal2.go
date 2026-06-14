package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

var Pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku(n int) {
	nPustaka = n
	for i := 0; i < nPustaka; i++ {
		fmt.Scan(&Pustaka[i].id, &Pustaka[i].judul, &Pustaka[i].penulis,&Pustaka[i].penerbit, &Pustaka[i].eksemplar, &Pustaka[i].tahun, &Pustaka[i].rating)
	}
}

func CetakTerfavorit(n int) {
	idx := 0
	for i := 1; i < n; i++ {
		if Pustaka[i].rating > Pustaka[idx].rating {
			idx = i
		}
	}
	fmt.Printf("%s, %s, %s, %d\n", Pustaka[idx].judul, Pustaka[idx].penulis, Pustaka[idx].penerbit, Pustaka[idx].tahun)
}

func UrutBuku(n int) {
	for i := 1; i < n; i++ {
		key := Pustaka[i]
		j := i - 1
		for j >= 0 && Pustaka[j].rating < key.rating {
			Pustaka[j+1] = Pustaka[j]
			j--
		}
		Pustaka[j+1] = key
	}
}

func Cetak5Terbaru(n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Println(Pustaka[i].judul)
	}
}

func CariBuku(n int, ratingCari int) {
	low, high := 0, n-1
	found := false
	for low <= high && !found {
		mid := (low + high) / 2
		if Pustaka[mid].rating == ratingCari {
			b := Pustaka[mid]
			fmt.Printf("%s, %s, %s, %d, %d, %d\n", b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
			found = true
		} else if Pustaka[mid].rating < ratingCari {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

	func main(){
	var n, ratingCari int
	fmt.Scan(&n)

	DaftarkanBuku(n)

	fmt.Println("Buku Terfavorit:")
	CetakTerfavorit(n)

	UrutBuku(n)

	fmt.Println("5 Buku Terbaru dengan Rating Tertinggi:")
	Cetak5Terbaru(n)

	fmt.Println("Masukkan rating buku yang ingin dicari:")
	fmt.Scan(&ratingCari)
	CariBuku(n, ratingCari)
	}
