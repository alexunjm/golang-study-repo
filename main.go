package main

func main() {
	cards := getCards()

	cards.print()
}

func getCards() deck {
	return deck{"Element1", "Element2"}
}