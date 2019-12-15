package main

import (
	"fmt"
	"runtime"
	"time"
)


func generate(ch chan int) {
	for i:= 2;;i++ {
		ch <- i
	}
}


func filter(in , out chan int, prime int) {
	for {
		i := <- in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main(){

	start := time.Now()

	timeout := time.After(4*time.Second)


	runtime.GOMAXPROCS(8)
	ch := make(chan int)
	go generate(ch)
	for {
		select {
			case prime := <-ch:
				fmt.Println(prime)
				ch1 := make(chan int)
				go filter(ch, ch1, prime)
				ch = ch1

			case <-timeout :
				fmt.Printf("\nexecution time %s", time.Since(start))
				return

		}
	}



}