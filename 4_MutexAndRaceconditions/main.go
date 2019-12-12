package main

import (
	"fmt"
	"sync"
	"time"
)



var i = 0

func worker(wg *sync.WaitGroup, m *sync.Mutex){

	// at this point is problem: If two threads execute i = i+1 at the same time, both has same initial value for i,
	// they ignores fact that one of thread would increment value of i first and second would start with incremented i.
	//thats why for large cyces proram returns < 1000, instead of 1000
	//Thats place where we use Mutex - which allows only one thread to access incrementing i (protected setion), so all threads comes out from value of previous thread run

	//waiting group must be here otherwise some of threads could wait in sleep but main thread will finish sooner so 1000 value will not be achieved at the end.
	m.Lock()
	i = i + 1
	m.Unlock()
	time.Sleep(10)
	wg.Done()
}


func main(){

	start := time.Now()


	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	for x := 0; x < 1000; x ++ {
		//wg.add must be within main thread, not in go routine.
		//the issue is that add will be called at the same time as wait which leads to runtime errors.
		wg.Add(1)
		go worker(&wg, &m)
	}

	wg.Wait()

	fmt.Printf("\nValue i is %v ",i)

	fmt.Printf("\nexecution time %s", time.Since(start))
}