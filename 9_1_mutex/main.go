package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)



func main(){

	start := time.Now()

	runtime.GOMAXPROCS(4)

	mutex := new(sync.Mutex)
	for i := 0; i <10; i++ {
		for j := 0 ; j < 10; j++ {
			//if I lock it here it is the same as synchronised app. Even slower , because only one thread could appear in gorutine and this new thread will unlock this
			// output is sorted btw as it ascts like synchronised
			mutex.Lock()
			go func(){
				fmt.Printf("%d + %d = %d \n", i, j , i+j)
				mutex.Unlock()
			}()
		}
	}

	fmt.Printf("\nexecution time %s", time.Since(start))
	fmt.Scanln()
}