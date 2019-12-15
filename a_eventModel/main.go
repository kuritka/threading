package main

import (
	"fmt"
	"time"
)



type Button struct {
	//evenet name is key of map, slice of channel string are listeners
	eventListeners map[string][]chan string
}

func NewButton() *Button {
	 result := new(Button)
	 result.eventListeners = make(map[string][]chan string)
	 return result
}

func (b *Button ) AddEventListener(event string, responseChannel chan string){
	if _, present := b.eventListeners[event];  present {
		b.eventListeners[event] = append(b.eventListeners[event], responseChannel)
	} else
	{
		b.eventListeners[event] = []chan string {responseChannel}
	}
}


func (b *Button ) RemoveEventListener(event string, listenerChannel chan string){
	if _, present := b.eventListeners[event];  present {
		for index, _ := range b.eventListeners[event] {
			if b.eventListeners[event][index] == listenerChannel {
				//remove element from array in GO!
				b.eventListeners[event] = append(b.eventListeners[event][:index], b.eventListeners[event][index+1:]...)
				break
			}
		}
	}
}

func (b *Button ) TriggerEvent(event string, data string) {
	if _, present := b.eventListeners[event];  present {
		for _, handler := range b.eventListeners[event] {
			go func(handler chan string){
				handler <- data
			}(handler)
		}
	}
}


func listenA(input <-chan string){
	for {
		msg := <- input
		fmt.Println("HELLO FROM A LISTENER, I'm updating database... ",msg)

	}
}


func listenB(input <-chan string){
	for {
		msg := <- input
		fmt.Println("HELLO FROM A LISTENER, I'm updating HDD... ",msg)

	}
}

func main(){

	start := time.Now()


	btn := NewButton()

	hOne := make(chan string)
	hTwo := make(chan string)

	btn.AddEventListener("click", hOne)
	btn.AddEventListener("click", hTwo)

	go listenA(hOne)
	go listenB(hTwo)

	fmt.Println("click event and two listeners")
	btn.TriggerEvent("click", "Clicked on XXX")

	//fmt.Println("remove first listener")
	btn.RemoveEventListener("click", hOne)
	btn.TriggerEvent("click", "Clicked on BUTTON ABAC")

	fmt.Scanln()

	fmt.Printf("\nexecution time %s", time.Since(start))
}