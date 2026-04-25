package main

import (
	"fmt"
	"math"
)

type titik struct {
	x	int
	y	int 
}

type lingkaran struct {
	koorPusat titik
	r		  int
}

func jarak(p, q titik) float64 {
	hasil := math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
	return hasil
}

func didalam(c lingkaran, p titik) bool {
	j := jarak(c.koorPusat, p)
	if j <= float64(c.r) {
		return true
	}
	return false
}

func main() {
	var l1, l2 lingkaran
	var t titik

	fmt.Scan(&l1.koorPusat.x, &l1.koorPusat.y, &l1.r)
	fmt.Scan(&l2.koorPusat.x, &l2.koorPusat.y, &l2.r)
	fmt.Scan(&t.x, &t.y)

	test1 := didalam(l1, t)
	test2 := didalam(l2, t)

	if test1 == true && test2 == true {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if test1 == true {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if test2 == true {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
