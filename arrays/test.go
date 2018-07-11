package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {
	fmt.Println("print test")
	text := "golang"
	buf := make([]byte, len(text))
	for i, c := range text {
		utf8.EncodeRune(buf, c)

		fmt.Println(i, reflect.TypeOf(c), string(buf), c)
	}

	rs := reverseString(text)
	fmt.Println(rs)
}

func reverseString(s string) []rune {
	// l := len(s)
	chars := []rune(s)
	// for i, c := range s {
	// 	chars[l-i-1] = c
	// }
	// return string(chars[2])
	return chars
}

// Implement an algorithm to determine if a string has all unique characters. What if you
// cannot use additional data structures?
func isUnique(text string) map[string]int {
	// iterate each character in string and map each ch to counter hash
	// for _, ch := range text {

	// }
	ch := make(map[string]int)
	return ch

}
