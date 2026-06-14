package main

import "fmt"


func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

func main() {
	var arr []int
	var n int = 0
	var input int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		arr = append(arr, input)
		n++
	}

	insertionSort(arr, n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	if n < 2 {
		fmt.Println("Data berjarak 0")
	} else {
		selisih := arr[1] - arr[0]
		berjarakSama := true

		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != selisih {
				berjarakSama = false
				break
			}
		}

		if berjarakSama {
			fmt.Printf("Data berjarak %d\n", selisih)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}