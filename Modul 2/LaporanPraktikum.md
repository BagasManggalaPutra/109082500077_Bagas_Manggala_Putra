# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Bagas Manggala Putra - 109082500077</p>

## Unguided 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul2/output/output1.png)
Jadi program ini gunanya untuk bisa dibilang menukar atau memindahkan nilai dari satu variabel ke variabel lain seperti output ditengah menjadi yang pertama output yang ada di pertama pindah ke output terakhir contoh "masa sih bang" menjadi "sih bang masa"


## Unguided 

### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):

```go
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

```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul2/output/output2.png)
Program yang ini dibuat untuk mengecek susunan dari warna tabung gelas reaksi dari gelas 1 sampai 4. nah jadi awalnya buat dulu variabel warna dari 1 sampai 4 tipe datanya string, buat juga variabel count yang di berikan nilai 0, habis itu kita buat perulangan dari 1 sampai 5 yang berisi inputan susunan warna dari 1 sampai 4 jika susunanya salah maka count akan bertambah 1 jika susunannya benar maka count tidak nambah, pas di akhir aku juga buat kondisi jika count tidak 0 maka outputnya berhasil: false dan jika count 0 maka berhasil : true

## Unguided 

### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):

```go
package main

import "fmt"

func main() {
	var brparsel, kg, sisagram, biayagram, biayakg, total,
		biayagramakhir int
	fmt.Print("Berat Parsel(gram): ")
	fmt.Scan(&brparsel)

	kg = brparsel / 1000
	sisagram = brparsel % 1000
	biayakg = kg * 10000

	if kg > 10 && sisagram < 500 {
		biayagram = 0
		biayagramakhir = sisagram * 15
	} else if kg > 10 && sisagram >= 500 {
		biayagram = 0
		biayagramakhir = sisagram * 5
	} else if sisagram < 500 {
		biayagram = sisagram * 15
		biayagramakhir = sisagram * 15
	} else {
		biayagram = sisagram * 5
		biayagramakhir = sisagram * 5
	}

	total = biayakg + biayagram
	
	fmt.Printf("Detail berat : %d kg + %d gr \n", kg, sisagram)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayakg, biayagramakhir)
	fmt.Printf("Total biaya : Rp. %d \n", total)
}


```
### Output Unguided : 

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/BagasManggalaPutra/109082500077_Bagas_Manggala_Putra/blob/main/Modul2/output/output3.png)
program ini dibuat untuk menghitung biaya pengiriman, awal kita buat variabel brparsel, kg, sisagram, biayagram, biayakg, total, biayagramakhir tipenya int, lanjut kita buat inputan untuk brparsel, nah habis itu buat rumus dimana untuk nilai dari kg diambil dari brparsel / 1000, sisagram didapat dari brparsel % 1000, dan biayakg diambil dari kg * 10000, habis itu buat logika kondisi 1 jika kg lebih besar dari 10 dan sisagram lebih kecil dari 500 biayagram jadi 0 dan biayagramakhir diambil dari sisagram dikali 15, kondisi 2 jika kg lebih besar dari 10 dan sisagram lebih besar atau sama dengan 500 maka biayagram jadi 0 dan biayagramakhir diambil dari sisagram dikali 5, kondisi 3 jika sisagram kurang dari 500 maka biayagram diambil dari sisagram dikali 15 dan biayagramakhir diambil dari sisagram dikali 15, kondisi terakhir biayagram akan diambil dari sisagram dikali 5 dan biayagramakhir diambil dari sisagram dikali 5, udah selesai untuk logika lanjut buat rumus untuk total biayakg ditambah biayagram yang terkhir aku buat output menggunakan fmt.Printf("Detail berat : %d kg + %d gr \n", kg, sisagram). fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayakg, biayagramakhir). fmt.Printf("Total biaya : Rp. %d \n", total)