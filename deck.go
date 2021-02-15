package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)	
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// return nil, err
		// log the error
		fmt.Println("Error:", err)
		// Option #1 - return a call to newDeck()
		// Option #2 - entirely quit the program
		os.Exit(1)
	}
	sliceString := strings.Split(string(byteSlice), ",")
	return deck(sliceString)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var newPosition int
	for i := range d {
		// newPosition := rand.Intn(len(d) -1)
		newPosition = r.Intn(len(d) -1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}