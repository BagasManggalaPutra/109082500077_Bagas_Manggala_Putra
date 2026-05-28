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