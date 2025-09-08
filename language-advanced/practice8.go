package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	wg.Add(1)
	go func(){
		for i:=0; i <= 100; i ++ {
			ch <- i
		}
		close(ch)
		defer wg.Done()
	}()
	
	wg.Add(1)
	go func(){
		for n := range ch {
			fmt.Printf("n-> %v", n)
		}
		defer wg.Done()
	}()

	wg.Wait()
}