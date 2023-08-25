package main

import (
	"fmt"
	 "runtime"
)

func main() {

	runtime.GOMAXPROCS(1)
	ch := make(chan struct{})
	go func() {
		fmt.Println("!!!!!child!!!!!!")
		ch <- struct{}{}
	}()
	for i := 100; i > 0; i-- {
		fmt.Printf("parent %d\n", i)
	}
	<-ch
}
