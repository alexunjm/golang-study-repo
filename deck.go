package main

import "fmt"

// Create a new type of 'deck'
// whic is a slice of strings
type deck []string


// every variable of type 'deck' can call this function
// on itself
// 'd' is a receiver (first letter as a convention)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}