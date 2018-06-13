package main

import (
	"fmt"
	"strings"
)

type Minutes int
type hours int
type Weight float64
type title string
type answer bool

func main() {

	// Example 1
	fmt.Println("\n---Convert underlying type to custom type---")
	minutes := Minutes(58)
	hours := hours(2)
	weight := Weight(945.7)
	name := title("the matrix")
	answer := answer(true)
	fmt.Println(minutes, hours, weight, name, answer)
	minutes += 3
	fmt.Println(minutes)

	// Example 2
	fmt.Println("\n---Compare values of custom type to its underlying type")
	if weight > Weight(907.3) {
		fmt.Println("Capacity exceeded")
	}
	if weight > 907.3 {
		fmt.Println("Capacity exceeded")
	}

	// Example 3: Methods
	fmt.Println("\n---Capitalize first letter of each word in title---")
	fixed := name.fixCase()
	fixed2 := title.fixCase(name) // hm, works the same as line above?
	fmt.Println(fixed)
	fmt.Println(fixed2)

	// Example 4
	fmt.Println("\n---Increment minutes---")
	for i := 1; i <= 3; i++ {
		minutes.Increment()
		// (&minutes).increment() // works the same as line above, but non-idiomatic
		fmt.Println(minutes)
	}
}

// strings.Title() capitalizes each word
func (t title) fixCase() string {
	return strings.Title(string(t)) // convert title to a string since type Title is based on a string.
}

// Increment receiver takes a pointer of type Minutes
func (m *Minutes) Increment() {
	*m = (*m + 1) % 60
}
