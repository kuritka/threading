package main

import (
	"fmt"
	"time"
)



func main(){

	start := time.Now()



	fmt.Printf("\nexecution time %s", time.Since(start))
}