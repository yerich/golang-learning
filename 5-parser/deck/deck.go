package deck

import (
	"fmt"
	"math/rand"
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
var cardValues = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

func (c Card) String() string {
	return fmt.Sprintf("%v%v", NUM_NAME[c.n], SUIT_NAME[c.suit])
}

func (c Card) N() int {
	return c.n
}

func (c Card) Suit() int {
	return c.suit
}

type Deck struct {
	cards []Card
}

func CardsValue(cards []Card) (int, int) {
	softValue, hardValue := 0, 0
	for _, card := range cards {
		if card.n == 1 {
			softValue += 1
			hardValue += 11
		} else {
			softValue += cardValues[card.n]
			hardValue += cardValues[card.n]
		}
	}
	return softValue, hardValue
}

func CreateDeck() *Deck {
	deck := new(Deck)
	deck.cards = make([]Card, 52)
	for i := 0; i < 52; i++ {
		deck.cards[i].n = i%13 + 1
		deck.cards[i].suit = i / 13
	}
	return deck
}

func (d Deck) RemoveCard(n int, suit int) bool {
	for i, c := range d.cards {
		if c.n == n && c.suit == suit {
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
			return true
		}
	}
	return false
}

func (d Deck) Cards() []Card {
	return d.cards
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
