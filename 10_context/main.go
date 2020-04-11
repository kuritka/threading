package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cancellation()
}

// cancellation: You are a developer and you hire a new robot. Your new
// robot knows immediately what they are expected to do and starts their
// work. You sit waiting for the result of the robot's work. The amount
// of time you wait on the robot is unknown because you need a
// guarantee that the result sent by the robot is received by you. Except
// you are not willing to wait forever for the robot to finish their work.
// They have a specified amount of time and if they are not done, you don't
// wait and walk away.
func cancellation() {

	ch := make(chan string)
	duration := time.Duration(2900) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	go func() {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		ch <- "Done"
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Expired")
	case <-ch:
		fmt.Println("Done")
	}

}
