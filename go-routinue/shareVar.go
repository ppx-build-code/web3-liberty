package main

import (
	"fmt"
	"time"
	"sync"
)

func testPall() {
	var counter int
	for i:=1; i<=10; i ++ {
		go func(){counter++}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("ct: %d", counter)
}

func testParallelChan() {
	done := make(chan int)
	ct := 1

	go func(){
		for delte := range done {
			ct += delte
		}
	}()

	for i:=1; i<=10; i ++ {
		go func() {
			done <- 1
		}()
	}

	fmt.Scanln()
    fmt.Println("Counter:", ct)
	
	
	// fmt.Println("ct: %d", ct)
}

func testParallelChan2() {
	var wg sync.WaitGroup
	ch := make(chan int)
	var counter int
	
	go func(){
		
		for c := range ch {
			counter += c
			
			// fmt.Println("compute, %d,%d", counter, c)
		}
	}()

	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println("send idx, %d", i)
			ch <- 1
			
		}()
	}

	wg.Wait()
	close(ch)
	time.Sleep(1 * time.Second)
	fmt.Println("Counter:", counter) // 1000


}

// 两数之和
// 方法一暴力双循环
func twoSum(nums [] int, target int) [2]int {
	for i,v := range nums {
		for j,z := range nums {
			if v + z == target {
				res:= [2]int{i,j}
				return res
			}
		} 
	}
	return [2]int{0,0}
}
// 空间换时间
func twoSum2(nums [] int, target int) []int {
	maps := make(map[int]int)
    result := make([]int,0)
	for i,v := range nums {
		exists:= target - v
		vi,ok := maps[exists]
		if ok {
            return append(result, vi, i)
			
		} else {
			maps[v] = i
		}
	}
	return result
}




func main() {
	// testPall()
	// testParallelChan()
	// testParallelChan2()
	nums := make([]int,0)
	nums = append(nums, 2,7,11,15)
	fmt.Printf("v:%v", twoSum(nums, 9))
	fmt.Printf("v:%v", twoSum2(nums, 9))
}
