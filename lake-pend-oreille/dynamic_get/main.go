package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		log.Fatal(err)
	}

	// reader.ReadAll() returns [][]string vs ioutil.ReadAll(),
	// which returns []string
	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// prints string
	fmt.Printf("%s", bs)
	// prints []byte
	fmt.Printf("%v", bs)
}
