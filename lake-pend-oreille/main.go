package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

// In Go, slices are built on top of arrays, and slices are dynamic, meaning they
// can grow and shrink. Whereas arrays are a fixed size.
func main() {
	// To fetch data dynamically:
	// res, err := http.Get("http://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2015.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// rdr := csv.NewReader(res.Body)
	// then implement the same code as below with rdr.

	// Open file and return type File
	f, err := os.Open("./Environmental_Data_Deep_Moor_2015.txt")
	// assume that a panic will be immediately fatal, for the entire program,
	// or at the very least for the current goroutine. Ask yourself "when this
	// happens, should the application immediately crash?" If yes, use a panic;
	// otherwise, use an error.
	if err != nil {
		panic(err)
	}
	// defer defers execution of fcn until surrounding fcn returns.
	// after you Open a file, make sure to Close the file.
	defer f.Close()
	fmt.Println(f) // prints a pointer to a memory address where the file is stored.

	// NewReader takes an io.Reader as an argument, io.Reader is an interface that
	// contains a Read method, and type File has the Read method, that means
	// File implicitly implements the Reader interface. This is polymorphism in Go.
	rdr := csv.NewReader(f)
	// reset delimiter to tab. Reader is a struct that has a Comma field.
	rdr.Comma = '\t'
	// TrimLeadingSpace ignores leading white spaces in a field.
	rdr.TrimLeadingSpace = true

	// ReadAll reads all remaining records from rdr, and returns [][]string because
	// each record (meaning each new line) is made up of multiple fields
	// (i.e. date, time, air_temp, etc).
	rows, err := rdr.ReadAll()
	// assume that a panic will be immediately fatal, for the entire program,
	// or at the very least for the current goroutine. Ask yourself "when this
	// happens, should the application immediately crash?" If yes, use a panic;
	// otherwise, use an error
	if err != nil {
		panic(err)
	}

	// loop each row
	for i, row := range rows {
		airTemp := row[1]
		baroPress := row[2]
		windSpeed := row[7]
		if i >= 1 && i < 5 {
			fmt.Println(i, row)
		}
		if i == 1 {
			// print type using %T format verb for desired values: air_temp,
			// barometric_press, and wind_speed fields.
			fmt.Printf("%T %T %T\n", airTemp, baroPress, windSpeed)
			fmt.Println(airTemp, baroPress, windSpeed)
			// break
		}
	}

	// calculate start time to process mean for 3 fields.
	start := time.Now()
	fmt.Println("total records: ", len(rows)-1) // -1 due to zero-based index
	fmt.Println("avg air temp: ", mean(rows, 1), median(rows, 1))
	fmt.Println("avg baro press: ", mean(rows, 2), median(rows, 2))
	fmt.Println("avg wind speed: ", mean(rows, 7), median(rows, 7))
	end := time.Now()
	// Sub returns the duration between two times.
	delta := end.Sub(start)
	fmt.Println("Abstracted mean performance: ", delta)

	median(rows, 1)
}

func mean(rows [][]string, idx int) float64 {
	var total, counter float64

	// loop each row
	for i, row := range rows {
		// skip header row
		if i != 0 {
			// ParseFloat converts string data from indexing row, to a float64
			// data type, which will later be used for calculations.
			val, _ := strconv.ParseFloat(row[idx], 64)
			// fmt.Printf("%T, %T %T\n", at, bp, ws)
			total += val
			counter++
		}
	}
	return total / counter // or: total / float64(len(rows)-1)
}

func median(rows [][]string, idx int) float64 {
	// holds data which will be sorted
	var sorted []float64

	for i, row := range rows {
		// skip header row
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			sorted = append(sorted, val)
		}
	}

	// fmt.Println("\nBEFORE SORTING")
	// for i, val := range sorted {
	// 	if i < 15 {
	// 		fmt.Printf("%f\n", val)
	// 	}
	// }
	// sort []float64 in increasing order
	sort.Float64s(sorted)

	// fmt.Println("\nAFTER SORTING")
	// for i, val := range sorted {
	// 	if i < 15 {
	// 		fmt.Printf("%f\n", val)
	// 	}
	// }

	if len(sorted)%2 == 0 { // meaning length is even
		mid := len(sorted) / 2
		higher := sorted[mid]
		lower := sorted[mid-1]
		return (higher + lower) / 2
	}
	mid := len(sorted) / 2
	return sorted[mid]
}
