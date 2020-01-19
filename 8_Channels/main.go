package main

import (
	"fmt"
	"time"
)

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage string
}

func main() {

	start := time.Now()

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To:      []string{"micha@xyz.com", "aaa@zzz.com"},
		From:    "aaa@aaa.com",
		Content: "Keep this in secret",
	}

	failed := FailedMessage{
		ErrorMessage:    "Some error",
		OriginalMessage: "Original Message",
	}

	msgCh <- msg
	errCh <- failed
	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case errorMsg := <-errCh:
		fmt.Println(errorMsg)
	//prevent from deadlock. it is important if no message in any channel otherwise deadlock runtime error
	default:
		fmt.Println("No essages received")
	}

	//
	//msgCh <- msg
	//errCh <- failed

	//fmt.Println(<-msgCh)
	//fmt.Println(<-errCh)

	fmt.Printf("\nexecution time %s", time.Since(start))

}
