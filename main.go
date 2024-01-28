package main

import (
	"fmt"
	"os"
)

func main() {
	deck, err := NewDeck(54)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(deck.toString())
}
