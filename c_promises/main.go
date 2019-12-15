package main

import (
	"ETL/utils"
	"errors"
	"fmt"
	"time"
)


/*
execution time 17.769µs
inserting  52fdfc07-2182-654f-163f-5f0f9a621d72
Failed to save purchase  Timeout occured
Second promise failed  Timeout occured
finish 52fdfc07-2182-654f-163f-5f0f9a621d72


execution time 23.781µs
Second promise failed  Timeout occured
Failed to save purchase  Timeout occured
inserting  52fdfc07-2182-654f-163f-5f0f9a621d72


execution time 13.261µs
inserting  52fdfc07-2182-654f-163f-5f0f9a621d72
finish 52fdfc07-2182-654f-163f-5f0f9a621d72
From PREMISE:  &{52fdfc07-2182-654f-163f-5f0f9a621d72 42.1 1}
Second promise succeed

 */

//promises, main thread is chaining functions synchronously but particular functions are running on their own threads
//also promisses adds add function handling fails.

type PurchaseOrder struct {
	Number string
	Value float64
	order int
}

type Promise struct {
	successChannel chan interface {}
	errorChannel chan error
}


func SavePOToDatabase(po *PurchaseOrder, shouldFail bool) *Promise {
	result := new(Promise)
	result.successChannel = make(chan interface{},1)
	result.errorChannel = make(chan error,1)


	go func() {
		//following line executes timeout, viz line 55
	//	time.Sleep(time.Second*2)
		if shouldFail {
			result.errorChannel <- errors.New("Failed to savepurchase order")
		} else {
			po.Number, _ = utils.Guid()
			fmt.Println("inserting ", po.Number)
			time.Sleep(1 * time.Second)
			fmt.Println("finish", po.Number)
			result.successChannel <- po
		}
	}()
	return result
}

//promise is defined
func (this *Promise) Then(success func(interface{}) error, failure func(error)) *Promise{
	result := new(Promise)

	//buffare must be setted at least to one because it could take sometime until someone drain the value
	result.successChannel = make(chan interface{},1)
	result.errorChannel = make(chan error,1)

	timeout := time.After(2*time.Second)
	go func(){
		select {
			case obj := <- this.successChannel:
				newErr := success(obj)
				if newErr == nil {
					result.successChannel <- obj
				} else {
					result.errorChannel <- newErr
				}
			case err := <- this.errorChannel :
				failure(err)
				result.errorChannel <-err
			case <- timeout :
				failure(errors.New("Timeout occured"))

		}
	}()

	return result
}


func main() {

	start := time.Now()

	po1 := new(PurchaseOrder)
	po1.Value = 42.1
	po1.order = 1

	SavePOToDatabase(po1,false).Then(func(obj interface{}) error{
			po := obj.(*PurchaseOrder)
			fmt.Println("From PREMISE: ",po)
			//second promise will respond on error if ERR will be returned here:
			//return errors.New("ERROR from FIRST PROMISE")
				return nil
	}, func(err error) {
			fmt.Println("Failed to save purchase ", err.Error())
	}).Then(func(interface{}) error{ fmt.Println("Second promise succeed"); return nil },
			func(err error){ fmt.Println("Second promise failed ", err.Error())})


	fmt.Printf("\nexecution time %s\n", time.Since(start))
	fmt.Scanln()
}
