package main

import (
	"fmt"
	"runtime"
	"time"
)



func main(){

	start := time.Now()
	//VARY CONFUSING WAY HOW TO IMPLEMENT CUSTOM MUTEX
	mutex := make(chan bool,1)
	runtime.GOMAXPROCS(4)

	for i := 0; i <10; i++ {
		for j := 0 ; j < 10; j++ {
			// fill one item to channel buffer and wait until it is flushed
			mutex<-true
			go func(){
				fmt.Printf("%d + %d = %d \n", i, j , i+j)
				//flushing
				<-mutex
			}()
		}
	}

	fmt.Printf("\nexecution time %s", time.Since(start))
	fmt.Scanln()
}