package main

import (
	"fmt"
	"sync"
)

func main() {

	shareVal := 10
	var mx sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i ++ {
		wg.Add(1)
		go func (){
			defer wg.Done()
			mx.Lock()
			shareVal ++
			mx.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("v:%v",shareVal)
}