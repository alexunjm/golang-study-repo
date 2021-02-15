package main


func sliceOfCards() []string {
	cards := []string{"Element 1", "Element 2"}
	cards = append(cards, "Element 3") // returns a new one
	return cards
}