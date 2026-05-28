# <h1 align="center">Laporan Praktikum Modul 14 - Selection sort </h1>

<p align="center">Bagas Manggala Putra - 109082500077</p>

## Unguided

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing-masing daerah.

```go
package main

import "fmt"

type rumahKerabat struct {
	nomor int
}

type arryrumahKerabat [1000000]rumahKerabat

func selectionSort(T *arryrumahKerabat, n int) {
	var i, j, idx_min int
	var t rumahKerabat
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if (*T)[j].nomor < (*T)[idx_min].nomor {
				idx_min = j
			}
			j++
		}
		t = (*T)[idx_min]
		(*T)[idx_min] = (*T)[i-1]
		(*T)[i-1] = t
		i++
	}
}

func main() {
	var T arryrumahKerabat
	var n, m int

	fmt.Print("Masukan banyak daerah: ")
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		fmt.Print("Masukan banyak rumah: ")
		fmt.Scan(&m)
		
		for j := 0; j < m; j++ {
			fmt.Scan(&T[j].nomor)
		}
		
		selectionSort(&T, m)
		
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Print(T[j].nomor)
			} else {
				fmt.Print(T[j].nomor, " ")
			}
		}
		fmt.Println()
	}
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul14/Output/output1.png)
Program ini digunakan buat ngurutin nomor rumah kerabat Hercules di tiap daerah dari nilai terkecil ke data terbesar / ascending pake algoritma Selection Sort. Di setiap daerah, program nerima input m buat nentuin jumlah rumah, terus baca semua nomor rumah ke array. Fungsi selectionSort itu kerjanya nyari elemen nilai terkecil dari posisi indeks saat ini sampe akhir array, terus nukerin sama elemen yang di posisi depan (i-1).

## Unguided

### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. Format Masukan masih persis sama seperti sebelumnya. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.

```go
package main

import "fmt"

type rumahKerabat2 struct {
	nomor int
}

type arryrumahKerabat2 [1000000]rumahKerabat2

func selectionSortasc2(T *arryrumahKerabat2, n int) {
	var i, j, idx_min int
	var t rumahKerabat2
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if (*T)[j].nomor < (*T)[idx_min].nomor {
				idx_min = j
			}
			j++
		}
		t = (*T)[idx_min]
		(*T)[idx_min] = (*T)[i-1]
		(*T)[i-1] = t
		i++
	}
}

func selectionSortdes2(T *arryrumahKerabat2, n int) {
	var t rumahKerabat2
	var i, j, idx_max int
	i = 1
	for i <= n-1 {
		idx_max = i - 1
		j = i
		for j < n {
			if (*T)[idx_max].nomor < (*T)[j].nomor {
				idx_max = j
			}
			j = j + 1
		}
		t = (*T)[idx_max]
		(*T)[idx_max] = (*T)[i-1]
		(*T)[i-1] = t
		i = i + 1
	}
}

func main() {
	var n int
	fmt.Print("Masukan banyak daerah: ")
	fmt.Scan(&n)

	for d := 0; d < n; d++ {
		var m int
		fmt.Print("Masukan banyak rumah: ")
		fmt.Scan(&m)

		var ganjil, genap arryrumahKerabat2
		nGanjil := 0
		nGenap := 0

		for k := 0; k < m; k++ {
			var x int
			fmt.Scan(&x)
			if x%2 != 0 {
				ganjil[nGanjil] = rumahKerabat2{nomor: x}
				nGanjil++
			} else {
				genap[nGenap] = rumahKerabat2{nomor: x}
				nGenap++
			}
		}

		selectionSortasc2(&ganjil, nGanjil)
		selectionSortdes2(&genap, nGenap)

		first := true
		for k := 0; k < nGanjil; k++ {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[k].nomor)
			first = false
		}
		for k := 0; k < nGenap; k++ {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(genap[k].nomor)
			first = false
		}
		fmt.Println()
	}
}


```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul14/Output/output2.png)
Program ini modifikasi dari soal pertama, yang di mana program bakal misahin nomor rumah ganjil dan genap ke dua array yang beda saat proses input data. Kalo seluruh data udah terinput dan terpisah ke kelompoknya masing-masing, proses pengurutan nya gini, untuk Nomor Rumah Ganjil: Diurutkan dari nilai terkecil ke terbesar / ascending pake fungsi selectionSortasc2. Proses pengurutan ke dua Nomor Rumah Genap: Diurutkan dari nilai terbesar ke terkecil (descending) pake fungsi selectionSortDesc. Kalo kedua kelompok array udah selesai diurutin, program langsung print nomor rumah ganjil dulu, terus dilanjutin sama nomor rumah genap di satu baris yang sama

## Unguided

### 3. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0. Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313. Keluaran adalah median yang diminta, satu data per baris.

```go

package main

import "fmt"
type dataBilangan struct {
	angka int
}

type arrData [1000000]dataBilangan

func selectionSort3(T *arrData, n int) {
	var i, j, idx_min int
	var t dataBilangan
	
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if (*T)[j].angka < (*T)[idx_min].angka {
				idx_min = j
			}
			j++
		}
		t = (*T)[idx_min]
		(*T)[idx_min] = (*T)[i-1]
		(*T)[i-1] = t
		i++
	}
}

func main() {
	var T arrData
	var jumlahData int
	var inputNum int

	for {
		_, err := fmt.Scan(&inputNum)
		if err != nil {
			break
		}

		if inputNum == -5313 {
			break
		}

		if inputNum == 0 {
			if jumlahData > 0 {
				selectionSort3(&T, jumlahData)

				if jumlahData%2 != 0 {
					medianIdx := jumlahData / 2
					fmt.Println(T[medianIdx].angka)
				} else {
					mid2 := jumlahData / 2
					mid1 := mid2 - 1
					median := (T[mid1].angka + T[mid2].angka) / 2
					fmt.Println(median)
				}
			}
		} else {
			if jumlahData < 1000000 {
				T[jumlahData].angka = inputNum
				jumlahData++
			}
		}
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul10/Output/output3.png)
Program ini dibuat untuk memantau dan menghitung nilai tengah (median) dari data yang dimasukin. Program bakal terus terima input angka berulang kali. Kalo angkanya bukan 0 dan bukan -5313, langsung disimpen ke array. Pas yang dimasukin angka 0, fungsi selectionSort3 bakal dipanggil buat ngurutin semua data yang udah masuk dari yang paling kecil ke paling besar. Setelah datanya urut, baru deh hitung mediannya semisal jumlah datanya ganjil, ambil langsung angka pas di tengahnya, kalo genap, diambil dua angka yang ada di tengah, dijumlahin, terus dibagi dua pakai pembagian bilangan bulat. Prosesnya baru berhenti total kalo masuk angka -5313

