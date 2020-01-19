package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	phrase := "The safest way to double your money is to fold it over and put it in your pocket\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}
	close(ch)

	for i := 0; i < len(words); i++ {
		fmt.Println(<-ch + " ")
	}
	fmt.Printf("\nexecution time %s", time.Since(start))

}
