package main

import (
	"encoding/json"
	"fmt"
)

// use these two custom types to demonstrate encoding and decoding
type resp struct {
	Page   int
	page   int
	Fruits []string
}

// uses tags during struct field declaration to customize JSON encoded key names
type resp2 struct {
	Page   int      `json:"page123"`
	Fruits []string `json: "fruits123"`
}

func main() {
	// first lets look at encoding basic data types to JSON strings
	boolB, _ := json.Marshal(true)
	fmt.Println(boolB, string(boolB))

	intB, _ := json.Marshal(4)
	fmt.Println(intB, string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(fltB, string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(strB, string(strB))

	// slices and maps encode to JSON arrays and objects
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(slcB, string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(mapB, string(mapB))

	// encode custom data types, using only exported fields in the encoded output
	res := &resp{
		Page:   1,
		page:   2,
		Fruits: []string{"apples", "oranges"},
	}
	r, _ := json.Marshal(res)
	fmt.Println(r, string(r))

	res2 := &resp2{
		Page:   2,
		Fruits: []string{"pears", "bananas"},
	}
	r2, _ := json.Marshal(res2)
	fmt.Println(r2, string(r2))

	// create a generic JSON string data structure to decode
	byt := []byte(`{"num": 6.13, "strs": ["a", "b"]}`)

	// variable to store the decoded data, which holds a map of strings to arbitrary data types
	var dat map[string]interface{}

	// actual decoding, and checks for associated errors
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// to use values in decoded map, cast them to the appropriate type. Here, cast the value in num to a float64 type
	num := dat["num"].(float64)
	fmt.Println(num)

	// accessing nested data requires a series of casts
	s := dat["strs"].([]interface{})
	str1 := s[0].(string)
	fmt.Println(s)
	fmt.Println(str1)
}
