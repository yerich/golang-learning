package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	DIAMONDS = iota
	CLUBS
	HEARTS
	SPADES
)

type Card struct {
	n    int
	suit int
}

var NUM_NAME = [...]string{"", "A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}
var SUIT_NAME = [...]string{"D", "C", "H", "S"}

func (c Card) String() string {
	return fmt.Sprintf("%v%v", NUM_NAME[c.n], SUIT_NAME[c.suit])
}

type Deck struct {
	cards []Card
}

func createDeck() *Deck {
	deck := new(Deck)
	deck.cards = make([]Card, 52)
	for i := 0; i < 52; i++ {
		deck.cards[i].n = i%13 + 1
		deck.cards[i].suit = i / 13
	}
	return deck
}

func (d Deck) Shuffle() {
	l := len(d.cards)
	for i := range d.cards {
		r := rand.Intn(l)
		d.cards[i], d.cards[r] = d.cards[r], d.cards[i]
	}
}

func (d Deck) String() string {
	var s string
	for i, c := range d.cards {
		if i > 0 {
			s += ","
		}
		s += c.String()
	}
	return s
}

func main() {
	rand.Seed(time.Now().UnixNano())
	d := createDeck()
	fmt.Printf("%v\n", d)
	d.Shuffle()
	fmt.Printf("%v\n", d)
}
