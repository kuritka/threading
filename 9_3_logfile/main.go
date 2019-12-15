package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)


const file = "./9_3_logfile/log.txt"
func main(){

	start := time.Now()

	f,_ := os.Create(file)
	f.Close()

	//large channel as disk is much slowed than CPU
	logCh := make(chan string, 50)

	go func () {
		for  msg := range logCh{
				f, _ := os.OpenFile(file, os.O_RDWR|os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				_, err := f.WriteString(logTime + " - " + msg)
				if err != nil {
					fmt.Println(err.Error())
				}
				f.Close()
		}
	}()
	//VARY CONFUSING WAY HOW TO IMPLEMENT CUSTOM MUTEX
	mutex := make(chan bool,1)
	runtime.GOMAXPROCS(4)

	for i := 0; i <10; i++ {
		for j := 0 ; j < 10; j++ {
			// fill one item to channel buffer and wait until it is flushed
			mutex<-true
			go func(){
				msg := fmt.Sprintf("%d + %d = %d \n", i, j , i+j)
				logCh<- msg
				fmt.Printf(msg)
				//flushing
				<-mutex
			}()
		}
	}

	fmt.Printf("\nexecution time %s", time.Since(start))
	fmt.Scanln()
}