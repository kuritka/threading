package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	//almost same as prvious example but having buffer
	ch := make(chan string,1)

	ch <- "Hello"

	fmt.Println(<-ch)

	fmt.Printf("\nexecution time %s", time.Since(start))

}
