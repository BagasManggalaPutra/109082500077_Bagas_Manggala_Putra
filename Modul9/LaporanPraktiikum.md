# <h1 align="center">Laporan Praktikum Modul 9 - Array </h1>

<p align="center">Bagas Manggala Putra - 109082500077</p>

## Unguided

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

```go
package main

import (
	"fmt"
	"math"
)

type titik struct {
	x	int
	y	int
}

type lingkaran struct {
	koorPusat titik
	r		  int
}

func jarak(p, q titik) float64 {
	hasil := math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
	return hasil
}

func didalam(c lingkaran, p titik) bool {
	j := jarak(c.koorPusat, p)
	if j <= float64(c.r) {
		return true
	}
	return false
}

func main() {
	var l1, l2 lingkaran
	var t titik

	fmt.Scan(&l1.koorPusat.x, &l1.koorPusat.y, &l1.r)
	fmt.Scan(&l2.koorPusat.x, &l2.koorPusat.y, &l2.r)
	fmt.Scan(&t.x, &t.y)

	test1 := didalam(l1, t)
	test2 := didalam(l2, t)

	if test1 == true && test2 == true {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if test1 == true {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if test2 == true {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul9/Output/output1.png)
Jadi program ini gunanya untuk menentukan posisi titik di 2 lingkaran. Diawal program menerima input dari user untuk koordinat titik dan radius dari 2 lingkaran, di program ini menggunakan 2 tipe bentukan yang pertama ada titik untuk menyimpan titik x dan y tipenya int dan ada tipe bentukan lingkaran untuk titik pusat yang tipenya titik dan r yang tipenya int. Ada dua fungsi utama yang dipake, yaitu fungsi jarak buat ngitung jarak antara dua titik rumusnya √((x1-x2)² + (y1-y2)²), dan fungsi didalam untuk ngecek apakah itu titik berada di dalam lingkaran, caranya bandingin jarak titik ke pusat lingkaran dengan radiusnya. Kalo jarak tersebut lebih kecil atau sama dengan radius, maka titik dinyatakan berada di dalam lingkaran. Setelah dua pengecekan itu, program bakal nampilin salah satu dari empat kemungkinan output yaitu titik berada di dalam lingkaran 1 dan 2, hanya di dalam lingkaran 1, hanya di dalam lingkaran 2, atau di luar keduanya.

## Unguided

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut: a. Menampilkan keseluruhan isi dari array. b. Menampilkan elemen-elemen array dengan indeks ganjil saja. c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap). d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna. e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil f. Menampilkan rata-rata dari bilangan yang ada di dalam array. g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut. h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.

```go
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

	//a Menampilkan keseluruhan isi dari array.
	semua(arr, n)

	//b Menampilkan elemen-elemen array dengan indeks ganjil saja.
	ganjil(arr, n)

	//c Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).
	genap(arr, n)

	//d Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.
	var x int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	kelipatanX(arr, n, x)

	//e Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
	var indx int
	fmt.Print("Masukkan indeks yang dihapus : ")
	fmt.Scan(&indx)
	hapusIndex(&arr, &n, indx)
	fmt.Print("Array setelah dihapus : ")
	semua(arr, n)

	//f Menampilkan rata-rata dari bilangan yang ada di dalam array.
	fmt.Printf("Rata-rata : %.2f\n", rata2(arr, n))

	//g Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.
	fmt.Printf("Standar deviasi : %.2f\n", standarDev(arr, n))

	//h Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.
	var mencari int
	fmt.Print("Cari frekuensi angka : ")
	fmt.Scan(&mencari)
	fmt.Println("Frekuensi", mencari, ":", frekuensi(arr, n, mencari), "kali")
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul9/Output/output2.png)
Program yang ini dibuat untuk menyimpan sekumpulan bilangan bulat menggunakan array, terus program ini akan melalukan banyak operasi ke array yang dibuat. pertama program minta input jumlah elemen yang ingin dimasukkan, terus minta input nilai tiap elemennya satu per satu. Kalao array udah ke isi, program melakukan beberapa hal berikut:

*Menampilkan semua elemen dilakukan dengan looping dari indeks 0 sampai n-1 lalu mencetak tiap elemennya.

*Menampilkan elemen indeks ganjil dan genap dilakukan dengan mengecek sisa bagi indeks terhadap 2. Jika i % 2 != 0 maka indeks tersebut ganjil, sebaliknya jika i % 2 == 0 maka genap.

*Menampilkan elemen indeks kelipatan x dilakukan dengan cara yang sama, yaitu mengecek apakah i % x == 0.

*Menghapus elemen pada indeks tertentu dilakukan dengan cara menggeser semua elemen yang ada di sebelah kanan indeks tersebut satu posisi ke kiri, kemudian nilai n dikurangi 1 sehingga elemen terakhir yang sudah tidak relevan tidak ikut ditampilkan.

*Menghitung rata-rata dilakukan dengan menjumlahkan semua elemen lalu dibagi dengan jumlah elemen n.

*Menghitung standar deviasi dilakukan dengan mencari selisih tiap elemen terhadap rata-rata, mengkuadratkannya, menjumlahkan semua hasil kuadrat tersebut, membaginya dengan n, lalu diakarkan menggunakan fungsi math.Sqrt.

*Menghitung frekuensi dilakukan dengan looping seluruh array dan menghitung berapa kali bilangan yang dicari muncul.



## Unguided

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

```go

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


```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul9/Output/output3.png)
program ini dibuat untuk menghitung seluruh hasil pertangdingan antara dua klub sepak bola. Diawali dengan program yang meminta input nama dari dua klub yang akan bertanding, kemudian meminta secara berulang input skor dari masing masing klub tiap pertandingan. ketika skor dimasukan program langsung akan menentukan siapa yang jadi pemenang di setiap pertandingan. kalo skor klubA lebih besar dari skor klubB maka nama dari klubA akan dimasukan ke array begitupun sebaliknya, nah jika skornya sama maka yang disimpan adalah "Draw". Proses input skor akan bakal jalan terus sampe salah satu atau kedua skor yang dimasukkan nilainya negatif, itu jadi tanda kalo pertandingannya udah selesai.

## Unguided

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.

```go

package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {

	var input rune
	*n = 0
	for {
		fmt.Scanf("%c", &input)
		if input == '.' || *n >= NMAX {
			break
		}
		if input != ' ' && input != '\n' {
			t[*n] = input
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {

	var tAsli tabel = t
	balikanArray(&t, n)

	for i := 0; i < n; i++ {
		if tAsli[i] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	fmt.Print("Teks\t\t: ")
	isiArray(&tab, &n)

	isPalindrom := palindrom(tab, n)

	fmt.Print("Reverse teks\t: ")
	balikanArray(&tab, n)
	cetakArray(tab, n)

	fmt.Printf("Palindrom\t? %t\n", isPalindrom)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul9/Output/output4.png)
program ini dibuat untuk menyimpan karakter yang diinput user, program ini menggunakan array bertipe rune buat nyimpen karakter tadi, inputnya dibaca satu persatu dan bakalan berhenti ketika ada inputan tanda titik (.). fungsi cetakArray gunanya untuk menampilkan semua karakter yang ada di array dengan tambahan spasi di setiap karakternya. Fungsi balikanArray pakai cara two-pointer, jadi ada dua penunjuk dari kiri dan kanan. nah dua elemen itu ditukar posisinya, terus yang kiri geser ke kanan dan yang kanan geser ke kiri sampai ketemu di tengah. fungsi palindrom ngecek apakah teks itu palindrom atau bukan dengan cara bandingin karakter dari depan dan belakang secara berpasangan. Kalau ada yang beda, berarti bukan palindrom. Tapi kalau semuanya sama, berarti itu palindrom dan fungsi bakal ngembaliin nilai true.
