package main

import "fmt"

func main() {
	cards := sliceOfCards()

	// fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
