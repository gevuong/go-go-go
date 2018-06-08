package main

import "fmt"

// EvenOdd : prints whether int is even or odd
func EvenOdd() {
	nums := []int{0, 1, 2, 3, 4}

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
