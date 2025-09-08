package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(1 * time.Second)
	fmt.Println("Working done...")
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	<- done
	fmt.Println("Done")
}