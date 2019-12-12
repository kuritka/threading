package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	ch := make(chan string)

	//here it fails in runtime as no consumer exists
	ch <- "Hello"

	//If you commment previous line it fails here as no producer exists
	fmt.Println(<-ch)

	fmt.Printf("\nexecution time %s", time.Since(start))

}
