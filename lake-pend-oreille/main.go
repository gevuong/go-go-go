package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// In Go, slices are built on top of arrays, and slices are dynamic, meaning they
// y can grow and shrink. Whereas arrays are a fixed size.
func main() {
	// Open file and return type File
	f, err := os.Open("./Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		panic(err)
	}
	// defer defers execution of fcn until surrounding fcn returns
	// So after you Open a file, make sure to Close the file.
	defer f.Close()
	fmt.Println(f) // returns a pointer to a memory address where the file is stored

	// NewReader takes an io.Reader as argument, and io.Reader is an interface, and a
	// File implements the Reader interface. This is how we do polymorphism in Go.
	rdr := csv.NewReader(f)
	// change delimiter to make comma a tab. type Reader is a struct has a Comma field.
	rdr.Comma = '\t'
	fmt.Println(rdr.TrimLeadingSpace) // default is set to false
	rdr.TrimLeadingSpace = true
	println(rdr.TrimLeadingSpace)
	// ReadAll reads all remaining records from rdr, and returns [][]string because
	// each record is made up of multiple fields.
	// Each record is basically one entry in a data file.
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		fmt.Println(i, row)
		if i == 1 {
			fmt.Println("%T %T %T\n", row[1], row[2], row[7])
			fmt.Println(row[1], row[2], row[7])
			break
		}
	}

	for i, row := range rows {
		if i != 0 && i < 10 {
			// ParseFloat parses/converts string data from indexing row,
			// to a float64 data type used for calculations
			at, _ := strconv.ParseFloat(row[1], 64)
			bp, _ := strconv.ParseFloat(row[2], 64)
			ws, _ := strconv.ParseFloat(row[7], 64)
			fmt.Printf("%T, %T %T\n", at, bp, ws)
		}
	}
}
