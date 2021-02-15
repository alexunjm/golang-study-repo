package main

func main() {
	cards := newDeck()

	// cards.print()
	// hold, change := deal(cards, 2)
	// hold, _ := deal(cards, 2)

	// hold.print()
	// change.print()

	// fmt.Println(hold.toString())

	// cards.saveToFile("my_cards")
	// newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}