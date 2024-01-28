package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	deck, err := NewDeck(36)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := rand.New(rand.NewSource(1))
	deck.Shuffle(r)
	deck.Print()

	hand1, hand2, err := Deal(deck, 6)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(hand1)
	fmt.Println(hand2)
}
