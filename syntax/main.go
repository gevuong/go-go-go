package main

import (
	"fmt"
	"go-linux/syntax/welcome"
	"log"
	"math"
	"net"
	"os"
	"reflect"
	"time"

	"github.com/golang/example/stringutil"
)

var myNumber = 1.25
var g = 1 // package-level variable

func main() {
	fmt.Println(stringutil.Reverse("hello"))
	roundUp := math.Ceil(myNumber)
	roundDown := math.Floor(myNumber)
	fmt.Println(roundUp, roundDown)

	fmt.Println(welcome.English)
	fmt.Println(welcome.Spanish)

	fmt.Println("-----Types-----")
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1.5))
	fmt.Println(reflect.TypeOf("hello"))
	fmt.Println(reflect.TypeOf(true))

	fmt.Println(reflect.TypeOf(net.IPv4(127, 0, 0, 1)))
	fmt.Println(reflect.TypeOf(time.Now()))

	fmt.Println("-----variable declarations-----")
	var a, b int // declare b and c variables of type int
	a, b = 1, 2
	var c, d = 3, 4
	e, f := 5, 6
	fmt.Println(a, b, c, d, e, f)

	fmt.Println("-----type conversions-----")
	var wholeNumber = 1
	var fractionNumber = 2.75
	var wholeNumber2 = int(fractionNumber)     // type conversion to int
	var fractionNumber2 = float64(wholeNumber) // type conversion to float64
	fmt.Println("wholeNumber2: ", wholeNumber2)
	fmt.Println("fractionNumber2: ", fractionNumber2)

	// convert to same type before using math or comparison operator
	fmt.Println(float64(wholeNumber) + fractionNumber)
	fmt.Println(float64(wholeNumber) < fractionNumber)

	fmt.Println("-----block scoping-----")
	var h = "h"
	{
		var i = "i"
		{
			var j = "j"
			fmt.Println(a, h, i, j)
		}
		// fmt.Println(a, h, i, j) // ERROR: "j" undefined
	}
	// fmt.Println(a, h, i, j) // ERROR: "i", "j" undefined

	fmt.Println("-----function declarations-----")
	fmt.Println(square(3))
	fmt.Println(add(2, 4))
	fmt.Println(subtract(3, 5))

	num, err := squareRoot(9)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)

	fmt.Println("-----return fileInfo-----")
	fileInfo, err := os.Stat("existent.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileInfo.Size())
	}

	fileInfo, err = os.Stat("nonexistent.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileInfo.Size())
	}
}

func square(num int) int {
	return num * num
}

func add(a int, b float64) (sum float64) {
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(float64(a) + b)
	fmt.Println(reflect.TypeOf(a))

	return float64(a) + b
}

// bare return
func subtract(a, b float64) (difference float64) {
	difference = a - b
	return
}

func squareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("cannot take square root of negative number")
	}
	return math.Sqrt(num), nil
}
