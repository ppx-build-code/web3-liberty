package main

import (
	"fmt"
)

func add(n *int) {
	*n = *n + 10
}

func main() {
	n := 1
	add(&n)
	fmt.Printf("number: %d", n)
}