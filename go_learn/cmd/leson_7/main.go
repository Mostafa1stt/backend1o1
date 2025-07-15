package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
// var m = sync.Mutex{}
var  wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main(){
	t0 := time.Now()
	for i:=0; i<1000; i++{
		wg.Add(1)
		go count(i)
	}
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	// fmt.Printf("the results are : %v\n",results)
}

func count(i int){

	var res int
	for i:=0 ; i<100000000;i++{
		res+=1
	}
	wg.Done()

}

// func dbCall(i int) {
// 	// Simulate DB call delay
// 	var delay float32 = 2000
// 	time.Sleep(time.Duration(delay)*time.Millisecond)
// 	// fmt.Println("The result from the database is:", dbData[i])
// 	// m.Lock()
// 	// save(dbData[i])	// results = append(results,dbData[i])
// 	// log()
// 	// m.Unlock()
// 	wg.Done()
// }
//
// func save(resualts string){
//
// 	m.Lock()
// 	results = append(results,resualts)
// 	m.Unlock()
// }
//
// func log(){
// 	m.RLock()
// 	fmt.Printf("the curent result are :%v\n",results)
// 	m.RUnlock()
// }
