package main

import (
	"fmt"
	"time"
	"sync"
)

func sayHello() {
	fmt.Printf("say hello!!!\n")
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Second)
	fmt.Printf("main func!!!")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
}

