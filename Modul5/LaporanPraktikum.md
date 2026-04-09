# <h1 align="center">Laporan Praktikum Modul 5 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Bagas Manggala Putra - 109082500077</p>

## Unguided 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	fmt.Printf("Deret Fibonacci suku ke-%d:\n", n)
	for i := 0; i <= n; i++ {
		fmt.Printf("Suku ke-%d = %d, ", i, fibonacci(i))
	}
	fmt.Println()
}

```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul5/Output/output1.png)
Jadi program ini gunanya untuk menghitung suku ke n dari deret bilangan fibonacci pake fungsi rekursif, deret ini di definisikan dengan suku ke 0 = 0, suku ke 1 = 1 dan suku berikutnya didapat dari penjumlahan dua suku sebelum. Fungsi fibonacci(n) bekerja dengan cara: jika n=0 kembalikan 0, jika n=1 kembalikan 1, selain itu kembalikan fibonacci(n-1) + fibonacci(n-2). Cara kerjanya user masukkin nilai N, terus program manggil fungsi fibonacci(N) baru nampilin hasilnya.


## Unguided 

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	pola(n)
}

func Bintang(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	Bintang(n - 1)
}

func pola(n int) {
	if n == 0 {
		return
	}

	pola(n - 1)
	Bintang(n)
	fmt.Println()
}

```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul5/Output/output2.png)
Program yang ini dibuat untuk nampilin pola segitiga bintang pake rekursif. User memberikan input n kemudian nilai n masuk ke fungsi pola di sini ada syarat jika n sama dengan 0 nilai akan di return, n di pola nantinya akan dikurang 1 terakhir kita membuat fmt.Println() untuk membuat baris baru, kemudian nilai n itu masuk ke fungsi bintang, di fungsi ini juga ada syarat jika n sama dengan 0 nilai akan di return kemudian membuat output fmt.Print("*") setelah itu nilai n di fungsi bintang dikurang 1.

## Unguided 

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja   yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
```go

package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	mencariFaktor(n, 1)
	fmt.Println() 
}

func mencariFaktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Printf("%d ", i)
	}

	mencariFaktor(n, i+1)
}


```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul5/Output/output3.png)
program ini dibuat untuk menampilkan semua faktor dari bilangan yang kita inputkan memakai fungsi mencariFaktor(n, i), didalam fungsi terdapat logika if jika i > n maka return artinya jika i sudah melebihi n fungsi akan diberhentikan, lalu ada juga logika if yang kedua jika n mod i sama dengan o maka akan menjalankan perintah di dalamnya, terakhir fungsi akan memanggil dirinya sendiri tapi nilai i ditambah 1 ini untuk program mengecek nilai i sampai ke n  