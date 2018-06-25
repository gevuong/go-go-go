package main

import "fmt"

type Car struct {
	Doors        int
	Transmission string
	Odometer     int
}

func main() {
	// Example 1
	var aValue = 1.23
	var aPointer = &aValue // add `&` to get memory address (or pointer) of aValue
	fmt.Println("\n---Example 1: pointers---")
	fmt.Println("aValue: ", aValue)
	fmt.Println("aPointer: ", aPointer)   // prints memory address of aValue
	fmt.Println("*aPointer: ", *aPointer) // add `*` in front of a pointer to turn pointer into value

	// Example 2
	fmt.Println("\n---Example 2: halving number---")
	myNumber := 2.6
	fmt.Println("myNumber before 'halve' is: ", myNumber)
	halve(&myNumber) // pass the pointer (or memory address) to halve() to modify original value of myNumber
	// halve(myNumber)
	fmt.Println("myNumber in 'main' is: ", myNumber)

	// Example 3
	fmt.Println("\n---Example 3: modify original value, car of type Car struct---")
	car := Car{
		Doors:        4,
		Transmission: "Automatic",
		Odometer:     36000,
	}

	fmt.Println(car.Doors, "doors")
	describe(&car)
	fmt.Println(car.Doors, "doors")

	// Example 4
	fmt.Println("\n---Example 4: array literal---")
	months := [3]string{"Jun", "Jul", "Aug"}
	costPerMonth := [3]float64{1.23, 2.34, 5.43}

	for i, month := range months {
		fmt.Println(month, costPerMonth[i])
	}

	// Example 5
	fmt.Println("\n---Example 5: slices")
	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]
	fmt.Println(a, s1, s2)

	s1 = s1[0:4]
	// s2 = s2[0:4] // slice bound out of range because we tried to extend slice beyond end of underlying array
	fmt.Println(a, s1, s2)

	// notice that cap(s1) > cap(s2). s2 doesn't have any room to grow because it's at the end of underlying array
	fmt.Println("len s1:", len(s1), "cap s1: ", cap(s1))
	fmt.Println("len s2:", len(s2), "cap s2: ", cap(s2))

	fmt.Println("\n---all slices are modified with changes to s2 and s1 because all slices point to same underlyling array---")
	s2[1] = 87
	s1[2] = 21
	fmt.Println(a, s1, s2)

	// create a new slice, that points to larger underlying array
	fmt.Println("\n---append creates new slice that points to larger underlying array---")
	s2 = append(s2, 5) // append creates a new underlying array that s2 now points to.
	fmt.Println(a, s1, s2)

	fmt.Println("\n---s2 slice is using a completely different underlying array than s1 slice---")
	s2[0] = 999
	fmt.Println(a, s1, s2)
	fmt.Println("len s1:", len(s1), "cap s1: ", cap(s1))
	fmt.Println("len s2:", len(s2), "cap s2: ", cap(s2)) // capacity doubles original capacity

	// Best practice to append elements to array slice
	fmt.Println("\n---best practice to append elements to slice---")
	s := []int{}
	s = append(s, 4)
	s = append(s, 6, 7)
	fmt.Println(s)

	// iterate each element in slice
	for i, v := range s {
		fmt.Println("element:", i, "value:", v)
	}

	// Example 6
	fmt.Println("\n---Example 5: map literal---")
	ages := map[string]int{"Alice": 9, "Bob": 12, "Carol": 18}
	fmt.Println(ages)
	for name, age := range ages {
		fmt.Println(name, age)
	}
}

//*float64 means to pass a pointer of type float64. *pointer turns pointer into original value
func halve(pointer *float64) {
	*pointer = *pointer / 2
	fmt.Println("myNumber in 'halve' is: ", *pointer)
}

// if you pass a pointer of type Car struct, you don't need to convert pointer to value using *c.Doors = 2?
func describe(c *Car) {
	c.Doors = 2
	fmt.Println(c.Doors, "doors")
	fmt.Println(c.Transmission, "car")
	fmt.Println(c.Odometer, "miles")
}
