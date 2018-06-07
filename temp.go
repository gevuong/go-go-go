package main

import (
	"fmt"
	"go-linux/welcome"
	"math"
)

var myNumber = 1.25

func main() {
	roundUp := math.Ceil(myNumber)
	roundDown := math.Floor(myNumber)

	fmt.Println(roundUp, roundDown)
	fmt.Println(welcome.English)
	fmt.Println(welcome.Spanish)
}
