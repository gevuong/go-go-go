package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamonds", "Clubs", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {

			// returns a new slice
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {

	for idx, card := range d {
		fmt.Println(idx, card)
	}
}
