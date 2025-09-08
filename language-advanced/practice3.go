package main

import (
	"fmt"
	"sync"
	"time"
)

func worker1(wg *sync.WaitGroup) {
	
	for i:=0; i<=10; i ++ {
		if i % 2 != 0 {	
			time.Sleep(10 * time.Millisecond)
			
			fmt.Println("worker1 print odd number: %d", i)
		}
		
	}
	defer wg.Done()
}

func worker2(wg *sync.WaitGroup) {
	for i:=0; i<=10; i ++ {
		if i % 2 == 0 {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("worker2 print even number: %d", i)
		}
	}
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker1(&wg)
	wg.Add(1)
	go worker2(&wg)
	wg.Wait()

}