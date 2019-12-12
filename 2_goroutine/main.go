package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main(){
 	start := time.Now()

 	/// * *  *  *   *  READ THIS *  *   *   *
 	//LIST of ID's is constantly changing!!! API is not immutable.
 	// Use ID's from http://dummy.restapiexample.com/api/v1/employees and example will work for another fiwe minutes
 	ids := []int {
 		1,
		114595,
		114617,
		114038,
		115714,
		117906,
	}

 //	numComplete := 0
	//concurency is ok here because during waiting for response , the thread can start new goroutine.
	//But try it ut with usage multiple cores ? (Now depends on response speed, if it is too slow one thread can handle. If it is fast, multiple cores can take care!)
	runtime.GOMAXPROCS(1)
 	wg := sync.WaitGroup{}

 	fmt.Printf("max core on your system: %v\n" , runtime.NumCPU())
 	for _, i := range ids {

		wg.Add(1)


		go func(id int,wg *sync.WaitGroup) {
			resp, _ := http.Get(fmt.Sprintf("http://dummy.restapiexample.com/api/v1/employee/%v", id))
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			empls := new(empl)
			err := json.Unmarshal(body, &empls)
			if err != nil {
				log.Panic("error " + strconv.Itoa(id))
			}
			fmt.Printf("%v\n", empls)
			//numComplete++
			wg.Done()
		}(i, &wg)
	}
 	//wait until all id's are proceed
 	wg.Wait()
	//for numComplete < len(ids) {
	//	time.Sleep(5 * time.Millisecond)
	//}
	elapsed := time.Since(start)
	fmt.Printf("\nexecution time %s", elapsed)
}


type empl struct {
	Id string  `json:"id"`
	Name string `json:"employee_name"`
	Salary string `json:"employee_salary"`
	Age string `json:"employee_age"`


}