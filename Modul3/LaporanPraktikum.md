# <h1 align="center">Laporan Praktikum Modul 3 - ALGORITMA DAN PEMROGRAMAN 2 Function </h1>
<p align="center">Bagas Manggala Putra - 109082500077</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n! (n−r)!, sedangkan C(n, r) = n! r!(n−r)!
#### soal1.go

```go

package main

import "fmt"

func main() {
	var a, b, c, d, p1, p2, c1, c2 int
	
	fmt.Print("Masukan Bilangan a b c d: ")
	fmt.Scanln(&a, &b, &c, &d)
	
	p1 = permutation(a, c)
	c1 = combination(a, c)
	
	p2 = permutation(b, d)
	c2 = combination(b, d)
	
	if a < c || b < d {
		fmt.Println("Input tidak valid harus a >= c dan b >= d")
	} else {
		fmt.Println(p1, c1)
		fmt.Println(p2, c2)
	}
}

func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutation(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func combination(n, r int) int{
		return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul3/Output/output1.png)
Program ini dibuat untuk menghitung kombinasi dan permutasi menggunakan function jadi rumus dari kombinasi dan permutasi dibuat didalam function sendiri, diawali membuat variabel di func main var a, b, c, d, p1, p2, c1, c2 tipe int, lanjut membuat code inputan dari variabel a, b, c, d, habis itu kita loncat dulu untuk membuat function baru yang bernama faktorial berparameter n tipenya int karena func ini mengembalikan nilai yang bertipe int, isinya var hasil int = 1, var i, kemudian kita membuat code perulangan for i = 1; i<=n; i++ { hasil = hasil * 1} kemudian di return hasil, habis dari func faktorial lanjut membuat function permutation kita buat dia memiliki 2 parameter n dan r yang bertipe int kemudian kita return yang berisi rumus dari permutasi, function terakhir yang kita buat adalah function combination yang memiliki 2 parameter n dan r tipenya int lalu kita return lagi yang isinya rumus dari kombinasi, balik ke function main kita membuat code pemanggilan function dari permutation dan combination lalu hasil dari pemanggilan tersebut disimpan ke dalam variabel dari p1, c1, p2, c2, yang terakhir membuat syarat menggunakan if else jika a lebih kecil dari c atau b lebih kecil dari d akan menghasilkan output "Input tidak valid harus a >= c dan b >= d" jika tidak akan langsung mengeluarkan output dari permutation dan combination.


## Unguided 

### 2. Diberikan tiga buah fungsi matematika yaitu f (x) = x 2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!

```go
package main

import "fmt"

func main() {
	var a, b, c, hasil1, hasil2, hasil3 int

	fmt.Print("Masukan Bilangan a b c: ")
	fmt.Scanln(&a, &b, &c)

	hasil1 = f(g(h(a)))
	hasil2 = g(h(f(b)))
	hasil3 = h(f(g(c)))

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}

func f(x int) int{
	return x * x
}

func g(x int) int{
	return x - 2
}

func h(x int) int{
	return x + 1
}

```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul3/Output/output2.png)
Program ini dibuat untuk menghitung fungsi dari matematika f(x) = xkuadrat, g(x) = x − 2 dan h(x) = x + 1. pada awal kita membuat var di func main a, b, c, hasil1, hasil2, hasil3 tipenya int dan membuat code inputan dari var a, b, c, skip dulu di func main lanjut membuat dari func f yang berparameter x tipenya int kita return karena xkuadrat jadi return x * x, lalu kita buat func lagi yaitu func g yang berparameter x tipe int kita return x - 2, func terakhir func h yang berparameter x tipe int kita return x + 1, balik ke func main kita buat rumus pemanggilan di dalam var hasil1, hasil2, hasil3, untuk hasil 1 (fogoh)(a) rumusnya f(g(h(a))), hasil2 (gohof)(b) rumusnya g(h(f(b))), hasil3 (hofog)(c) rumusnya h(f(g(c))) kemudian kita membuat code untuk output menggunakan fmt.Println yang berisi var hasil1, hasil2, hasil3

## Unguided 

### 3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	in1 := didalam(cx1, cy1, r1, x, y)
	in2 := didalam(cx2, cy2, r2, x, y)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(x, y, cx, cy) <= r
}

```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul3/Output/output3.png)
Program ini dibuat untuk mendefinisikan suatu lingkaran menggunakan titik koordinat pusat (cx, cy) dengan radius r, diawal code dalam import kita tidak hanya menggunakan fmt tetapi kita juga menggunakan math untuk fungsi akar kuadrat, nah di func main kita membuat var dari cx1, cy1, r1 tipe float64, var cx2, cy2, r2 tipe float64, var x, y tipe float64, lanjut membuat code inputan untuk cx1, cy1, r1, cx2, cy2, r2, x, y, seperti 2 program di atas kita loncat dulu dari func main sekarang kita buat dulu func baru yang namanya func jarak parameternya a, b, c, d tipenya float64 kita return menggunakan rumus dari jarak akar kuadrat dari (a-c)kuadrat + (b-d)kuadrat dalam code math.Sqrt((a-c)*(a-c) + (b-d)*(b-d)), lanjut membuat func lagi 1 yang namanya func didalam mempunyai parameter cx, cy, r, x, y tipenya float64 tetapi func ini mengembalikan nilai true atau false karena memakai bool, kemudian kita return jarak(x, y, cx, cy) <= r, lanjut ke func main lagi kita membuat variabel in1 dan in2 untuk memanggil hasil dari func didalam, lalu kita membuat logika if jika in1 dan in2 maka "Titik di dalam lingkaran 1 dan 2", jika in1 maka "Titik di dalam lingkaran 1", jika in2 maka "Titik di dalam lingkaran 2" jika tidak ada syarat yang terpenuhi maka "Titik di luar lingkaran 1 dan 2"