package main

import (
	"ETL/utils"
	"fmt"
	"time"
)


/*
call backs are totaly asynchronous stuff...

inserting  88332683-3f63-b566-fab7-600edbccde40
inserting  37aae9fa-1752-3a55-7d4b-953b2542dea1
finish 37aae9fa-1752-3a55-7d4b-953b2542dea1
finish 88332683-3f63-b566-fab7-600edbccde40
PO:  &{37aae9fa-1752-3a55-7d4b-953b2542dea1 -2 2}
PO:  &{88332683-3f63-b566-fab7-600edbccde40 42.27 1}

execution time 2.000385979s


 */

type PurchaseOrder struct {
	Number string
	Value float64
	order int
}


//some long term asynchronous task...
func SavePOToDatabase(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number,_ = utils.Guid()
	fmt.Println("inserting ",po.Number)
	time.Sleep(2*time.Second)
	fmt.Println("finish", po.Number)
	callback <- po
}

func main() {

	start := time.Now()

	po1 := new(PurchaseOrder)
	po2 := new(PurchaseOrder)

	po1.Value = 42.27
	po2.Value = -2
	po1.order = 1
	po2.order = 2

	ch := make(chan *PurchaseOrder)

	go SavePOToDatabase(po1, ch)

	go SavePOToDatabase(po2, ch)

	newPo := <- ch

	fmt.Println("PO: ", newPo)

	newPo = <- ch

	fmt.Println("PO: ", newPo)



	fmt.Printf("\nexecution time %s", time.Since(start))

}
