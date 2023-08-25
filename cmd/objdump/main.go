package main

import (
	"fmt"
)

func xorit(s1, s2, s3 []int) {
	for i := 99; i > 0; i-- {
		s3[i] = s1[i] ^ s2[i]
	}
}

func main() {
	slice1 := make([]int, 100)
	slice2 := make([]int, 100)
	slice3 := make([]int, 100)
	for i := 99; i >= 0; i-- {
		slice1[i] = i
		slice2[i] = 100 - i
	}
	xorit(slice1, slice2, slice3)
	for i := 99; i > 0; i-- {
		fmt.Printf("parent %d - %d\n", i, slice3[i])
	}
}
