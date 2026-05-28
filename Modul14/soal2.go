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
