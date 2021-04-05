package main

import (
	"fmt"
	"math/rand"
	"time"

	"gotest/3-bust/deck"
)

var cardValues = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

const HIT_SOFT_17 bool = false

func calcBust(tries int, startCardValue int) int {
	busts := 0

	for i := 0; i < tries; i++ {
		d := deck.CreateDeck()
		d.Shuffle()

		cards := d.Cards()

		var softValue, hardValue int
		if startCardValue == 1 {
			softValue, hardValue = 11, 1
		} else {
			softValue, hardValue = cardValues[startCardValue], cardValues[startCardValue]
		}
		// fmt.Printf("======\n")

		for i := 0; i < len(cards); i++ {
			if (!HIT_SOFT_17 && softValue >= 17 || softValue >= 18) && (hardValue >= 17 || softValue <= 21) {
				break
			}

			if cards[i].N() == 1 {
				hardValue += 1
				if softValue > hardValue {
					softValue += 1
				} else {
					softValue += 11
				}
			} else {
				softValue += cardValues[cards[i].N()]
				hardValue += cardValues[cards[i].N()]
			}
			// fmt.Printf("%v %d %d\n", cards[i], softValue, hardValue)
		}
		if hardValue >= 22 {
			// fmt.Printf("bust\n")
			busts++
		}
	}

	return busts
}

func main() {
	rand.Seed(time.Now().UnixNano())

	tries := 10000000000

	for i := 0; i <= 0; i++ {
		busts := calcBust(tries, i)
		fmt.Printf("%v: %d tries, %d busts\n", deck.NUM_NAME[i], tries, busts)
	}
}
