# <h1 align="center">Laporan Praktikum Modul 4 - ALGORITMA DAN PEMROGRAMAN 2 Function </h1>
<p align="center">Bagas Manggala Putra - 109082500077</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d.Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n! (n−r)! , sedangkan C(n, r) = n! r!(n−r)!
#### soal1.go

```go

package main

import "fmt"

func main() {
	var a, b, c, d int
	var p1, p2, c1, c2 int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &p1)
	combination(a, c, &c1)
	fmt.Println(p1, c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)
	fmt.Println(p2, c2)
}

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n int, r int, hasil *int) {
	var fakN, fakNR int
	
	factorial(n, &fakN)
	factorial(n-r, &fakNR)   
	
	*hasil = fakN / fakNR
}
func combination(n int, r int, hasil *int) {
	var fakN, fakR, fakNR int
	
	factorial(n, &fakN)
	factorial(r, &fakR)      
	factorial(n-r, &fakNR)   
	
	*hasil = fakN / (fakR * fakNR)
}

```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul4/Output/output1.png)
Program ini dibuat untuk mencari hasil dari permutasi dan kombinasi menggunakan rumus P(n, r) = n! (n−r)! , dan C(n, r) = n! r!(n−r)!, kita juga menggunakan pointer (*) agar hasilnya bisa langsung di simpan dalam alamat variabel dan juga untuk mengembalikan nilai, program ini mempunyai 4 variabel inputan a, b, c, d inputan ini di gunakan untuk permutasi dan kombinasi (a, c) dan (b, d). Didalam program juga ada perulangan untuk mencari faktorial, nah prosedur permutasi dan kombinasi akan terus manggil prosedur faktorial untuk menghitung hasilnya menggunakan rumus di atas tadi

## Unguided 

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal.Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer) Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan. Keterangan: Astuti menyelesaikan 6 soal dalam waktu 287 menit, sedangkan Bertha 7 soal dalam waktu 294 menit. Karena Bertha menyelesaikan lebih banyak, maka Bertha menang. Jika keduanya menyelesaikan sama banyak, maka pemenang adalah yang menyelesaikan dengan total waktu paling kecil.

```go

package main

import "fmt"

func main() {
	var nama, pemenang string
	var soal, skor, maxsoal, minskor int
	maxsoal = -1
	minskor = 0
	adapeserta := false

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		hitungskor(&soal, &skor)
		adapeserta = true

		if soal > maxsoal || (soal == maxsoal && skor < minskor) {
			maxsoal = soal
			minskor = skor
			pemenang = nama
		}
	}

	if adapeserta {
		fmt.Println(pemenang, maxsoal, minskor)
	} else {
		fmt.Println("Tidak ada peserta")
	}

}
func hitungskor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu, i int
	for i = 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul4/Output/output2.png)
Program ini dibuat untuk menentukan pemenang lomba pemograman tingkat nasional inputannya ada nama peserta dan 8 nilai waktu dalam menit. kalo ngerjain lebih dari 301 menit dianggap gagal. Dalem program pake prosedur hitungskor gunanya untuk ngebaca 8 waktu, ngitung soal yang berhasil kurang dari 301 menit dan total waktu dari soal yang berhasil. Ini akan berulang sampai ada kata "Selesai". Pemenang dipilih dari jumlah soal paling banyak dan total waktu tercepat. hasil yang menjadi output ada nama pemenang, jumlah soal, dan total waktu