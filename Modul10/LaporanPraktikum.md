# <h1 align="center">Laporan Praktikum Modul 10 - Pencarian Nilai Max/Min </h1>

<p align="center">Bagas Manggala Putra - 109082500077</p>

## Unguided

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

```go
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

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul10/Output/output1.png)
Jadi program dibuat untuk mendata berat anak kelinci paling kecil dan paling besar berdasarkan input dari user, diawal program minta user masukin jumlah kelinci yang mau ditimbang, terus user diminta untuk masukin berat setiap kelinci satu persatu yang datanya itu dimasukin ke array. lanjut programnya nyari nilai terkecil dan terbesar. Terus program loop dari data kedua sampai terakhir. Tiap kali nemu berat yang lebih kecil dari patokan terkecil, patokan itu diganti, begitu juga kalo nemu yang lebih besar dari patokan terbesar, patokan terbesarnya diganti Setelah semua data dicek, program nampilin hasil berat kelinci paling kecil dan paling besar.

## Unguided

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

```go
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


```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul10/Output/output2.png)
Program yang ini dibuat untuk membagi berat ikan ke beberapa wadah, kita membuat inputan x untuk jumlah ikan dan y untuk kapasitas wadah, serta array untuk menyimpan berat masing masing ikan. Selanjutny digunakan nestedloop untuk menjumlahkan berat ikan ke dalam setiap wadah sesuai kapasitas y yang ditentukan. Di dalam perulangan tersebut, selanjutnya total berat tiap wadah ditampilkan secara berurutan, lalu dihitung juga total berat keseluruhan dan jumlah wadahnya. Terakhir program menghitung rata-rata berat per wadah dengan membagi total berat semua ikan dengan jumlah wadah yang digunakan.

## Unguided

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

```go

package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var data arrBalita
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(data, n))
}


```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul10/Output/output3.png)
program ini dibuat untuk mencari berat balita paling kecil, paling besar, dan rata rata beratnya. pertama tama user diminta untuk masukin jumlah balita terus masukin berat badan setiap balita satu persatu ke array, ada 2 fungsi yang dipake, pertama fungsi untuk mencari nilai max dan min, fungsi kedua untuk mencari rata rata caranya jumlahin semua data terus dibagi n. Kalo udah dapet semua hasil, program nampilin tiga output: berat terkecil, berat terbesar, dan rata-rata.

