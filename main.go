package main

import (
	"fmt"
	"os"
)

func main() {
	deck, err := NewDeck(52)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	deck.Print()

	hand1, hand2, err := Deal(deck, 52)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(hand1)
	fmt.Println(hand2)
}
