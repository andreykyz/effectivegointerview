package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "世界"
	fmt.Printf("Hello, %s %d %d %d\n", str, len(str), len([]rune(str)), utf8.RuneCountInString("世界"))
	for i, ch := range str {
		fmt.Printf("%s %d \n", string(ch), i)
	}
}
