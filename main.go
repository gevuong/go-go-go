package main

import (
	"fmt"
	"go-linux/welcome"
	"math"

	"github.com/golang/example/stringutil"
)

var myNumber = 1.25

func main() {
	fmt.Println(stringutil.Reverse("hello"))
	roundUp := math.Ceil(myNumber)
	roundDown := math.Floor(myNumber)

	fmt.Println(roundUp, roundDown)
	fmt.Println(welcome.English)
	fmt.Println(welcome.Spanish)
}
