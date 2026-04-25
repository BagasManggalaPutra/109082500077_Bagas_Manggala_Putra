package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {

	var input rune
	*n = 0
	for {
		fmt.Scanf("%c", &input)
		if input == '.' || *n >= NMAX {
			break
		}
		if input != ' ' && input != '\n' {
			t[*n] = input
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	
	var tAsli tabel = t
	balikanArray(&t, n)

	for i := 0; i < n; i++ {
		if tAsli[i] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	fmt.Print("Teks\t\t: ")
	isiArray(&tab, &n)

	isPalindrom := palindrom(tab, n)

	fmt.Print("Reverse teks\t: ")
	balikanArray(&tab, n)
	cetakArray(tab, n)

	fmt.Printf("Palindrom\t? %t\n", isPalindrom)
}