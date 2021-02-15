package main

import "fmt"

// Create a new type of 'deck'
// whic is a slice of strings
type deck []string


func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
		
	}

	return cards
}


// every variable of type 'deck' can call this function
// on itself
// 'd' is a receiver (first letter as a convention)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}