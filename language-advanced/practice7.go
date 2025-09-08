
package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	done := make(chan int)
	wg.Add(1)
	go func (){
		for i := 1; i<=10; i++ {
			done <- i
		}
		defer wg.Done()

	}()
	go func() {
		for n := range done {
			fmt.Printf("value:%d\n",n)
		}
		
	}()
	wg.Wait()
	close(done)
	
	fmt.Printf("over.")
}