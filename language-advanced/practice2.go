package main

import (
	"fmt"
)

func sliceNum(nums []int) {
	for i := range nums {
		nums[i] = nums[i] * 2
	}
}

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1,2,3,4,5,6)
	sliceNum(nums)
	for _,v := range nums {
		fmt.Println("v:%d", v)
	}
}