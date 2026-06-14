# <h1 align="center">Laporan Praktikum Modul 14 - Insertion sort </h1>

<p align="center">Bagas Manggala Putra - 109082500077</p>

## Unguided

### 1.Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".

```go

package main

import "fmt"


func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

func main() {
	var arr []int
	var n int = 0
	var input int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		arr = append(arr, input)
		n++
	}

	insertionSort(arr, n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	if n < 2 {
		fmt.Println("Data berjarak 0")
	} else {
		selisih := arr[1] - arr[0]
		berjarakSama := true

		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != selisih {
				berjarakSama = false
				break
			}
		}

		if berjarakSama {
			fmt.Printf("Data berjarak %d\n", selisih)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul14/Output/output1.png)
Program ini buat membaca kumpulan angka yang bukan negatif, diurutkan pake insertion sort, dan ngecek jarak antara angka yang udah diurutin apakah jaraknya sama apa nggak, pertama buat func insertionsort Tiap elemen ke-i disimpen sementara di temp, terus dibandinginn sama elemen-elemen di sebelah kirinya. Kalau elemen kiri lebih besar dari temp, elemen itu digeser ke kanan. Kalo udah ketemu elemen yang lebih kecil atau udah mentok ke ujung kiri, temp langsung dimasukin ke posisi kosong yang masih sisa. Di func main membuat perulangan for untuk menampung angka dari pengguna. Angka dimasukkan ke dalam slice (arr). Perulangan ini baru akan berhenti kalau memasukkan angka negatif. Jadi, angka negatif di sini berfungsi sebagai tanda perulangan berhenti. Kemudian untuk mencari selisih digunakan arr[1] - arr[0]. Terus dicek satu persatu apakah selisih antar elemen berikutnya selalu sama. Kalo ada yang beda, langsung jadi false dan loop berhenti. Kalo sampe habis semuanya sama, berarti datanya berjarak tetap dan nilai jaraknya diprint.

## Unguided

### 2. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini: const nMax : integer = 7919 type Buku = < id, judul, penulis, penerbit : string eksemplar, tahun, rating : integer > type DaftarBuku = array [ 1..nMax] of Buku Pustaka : DaftarBuku nPustaka: integer Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari. Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

```go
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


```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul14/Output/output2.png)
Program ini guna untuk mengelola data buku di perpustakaan. Ada 4 fungsi utama yaitu DaftarkanBuku, CetakTerfavorit, UrutBuku, Cetak5Terbaru. Func DaftarkanBuku bakal baca jumlah buku (n), terus minta kamu masukin detail tiap buku satu per satu (ID, judul, penulis, dll) buat disimpan ke dalam variabel pustaka. func CetakTerfavorit Dia bakal mencari semua buku yang udah didaftarin, bandingin ratingnya satu sama lain, terus kalo pas ketemu yang ratingnya paling tinggi, dia langsung cetak detail buku itu ke layar. Func UrutBuku ini pakai metode Insertion Sort buat ngurutin semua buku berdasarkan ratingnya dari yang paling tinggi ke yang paling rendah. Nah pas selesai buku-buku dengan rating gede bakal otomatis ada di paling atas. func Cetak5Terbaru nah pas urutannya udah rapi karena  fungsi UrutBuku, fungsi ini tinggal ambil 5 buku pertama di barisan paling depan. func CariBuku ini fitur digunakan untuk nyari buku berdasarkan rating yang dicari. Ini pakai teknik Binary Search supaya pencariannya cepat karena datanya udah diurutin sama UrutBuku. fungsi ini bakal terus mencari sampai ketemu buku yang ratingnya sesuai sama yang mau dicari.

## Unguided