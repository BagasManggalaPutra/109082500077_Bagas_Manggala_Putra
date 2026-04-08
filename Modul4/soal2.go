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
