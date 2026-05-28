package main

import "fmt"

type arryInt [1000000]int

func selectionSort3(T *arryInt, n int) {
	var i, j, idx_min int
	var t int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if (*T)[j] < (*T)[idx_min] {
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
	var T arryInt
	var input int
	count := 0

	for {
		fmt.Scan(&input)

		if input == -5313 {
			break
		}

		if input == 0 {
			if count > 0 {
				selectionSort3(&T, count)

				if count%2 != 0 {
					median := T[count/2]
					fmt.Println(median)
				} else {
					mid1 := T[(count/2) - 1]
					mid2 := T[count/2]
					median := (mid1 + mid2) / 2
					fmt.Println(median)
				}
			}
		} else {
			T[count] = input
			count++
		}
	}
}