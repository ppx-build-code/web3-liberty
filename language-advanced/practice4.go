package main

import (
	"fmt"
	"time"
)

func dispatch(funcs [] func()) {
	for i := range funcs {
		go func (){
			start := time.Now()
			funcs[i]()
			end := time.Now()
			fmt.Printf("任务 %v 执行耗时 %v \n", i, end.Sub(start).Milliseconds())
		}()
	}
}

func main() {
	funcs := make([]func(), 0)
	first := func(){
		fmt.Println("first")
		time.Sleep(10 * time.Millisecond)
	}
	second := func(){
		fmt.Println("second")
		time.Sleep(10 * time.Millisecond)
	}
	third := func(){
		fmt.Println("third")
		time.Sleep(10 * time.Millisecond)
	}
	funcs = append(funcs, first, second, third )
	dispatch(funcs)
	time.Sleep(3*time.Second)
}