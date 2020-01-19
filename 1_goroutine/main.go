package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	godur, _ := time.ParseDuration("10ms")
	//allows to use N processors for its execution
	runtime.GOMAXPROCS(2)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
			time.Sleep(godur)
		}
	}()

	func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(godur)
		}
	}()

	time.Sleep(1 * time.Second)
}
