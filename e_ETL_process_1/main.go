package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)


type Product struct {
	PartNumber string
	UnitCost float64
	UnitPrice float64
}


type Order struct {
	CustomerNumber int
	PartNumber string
	Quantity int
	UnitCost float64
	UnitPrice float64
}

func extract(ch chan *Order) {
	f,_ := os.Open("./e_ETL_process_1/input/orders.txt")
	defer f.Close()

	r := csv.NewReader(f)
	//interesting!
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber,_ = strconv.Atoi(record[1])
		order.PartNumber = record[0]
		order.Quantity,_ = strconv.Atoi(record[2])
		ch <- order
	}
	close(ch)
}


 func transform(extractChannel , transformChannel chan *Order ) {
	f,_ := os.Open("./e_ETL_process_1/input/orders.txt")
	//https://www.joeshaw.org/dont-defer-close-on-writable-files/
	 defer f.Close()

	 r := csv.NewReader(f)
	 records,_ := r.ReadAll()
	 productList := make(map[string]*Product)
	 for _, record := range records {
	 	product := new(Product)
	 	product.PartNumber = record[0]
	 	product.UnitCost, _= strconv.ParseFloat(record[1], 64)
	 	product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.PartNumber] = product
	 }

	 for  o := range extractChannel {
	 	//web service call
	 	time.Sleep(4* time.Millisecond)
	 	o.UnitPrice = productList[o.PartNumber].UnitPrice
	 	o.UnitCost = productList[o.PartNumber].UnitCost
	 	transformChannel <- o
	 }
	close(transformChannel)
 }

func load(transformChannel chan *Order, done chan bool)  {

	f,_ := os.Open("./e_ETL_process_1/output/dest.txt")

	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n","Part Nuumber", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")

	for o := range transformChannel{
		time.Sleep(1* time.Millisecond)
		fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s",o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice, o.UnitCost * float64(o.Quantity), o.UnitPrice* float64(o.Quantity))
	}

	_ = f.Close()
	//if err != nil {
	//	panic(err.Error())
	//}
	done <- true
}



func main(){

	fmt.Println("started...")
	start := time.Now()

	extractCh := make (chan *Order)
	transformCh := make (chan *Order)
	doneCh := make(chan bool)

	go extract(extractCh)
	go transform(extractCh, transformCh)
	go load(transformCh, doneCh)

	<- doneCh

	fmt.Printf("\nexecution time %s", time.Since(start))
}