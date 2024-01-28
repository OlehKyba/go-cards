package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	deck, err := NewDeck(36)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	deck.Shuffle(r)
	fmt.Println("Generated deck: ", deck.toString())

	hand1, hand2, err := Deal(deck, 6)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("After deal:")
	fmt.Println("First hand: ", hand1.toString())
	fmt.Println("Second hand: ", hand2.toString())

	err = SaveDeckToFile("./dumped_decks/main_deck.txt", deck)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		err = RemoveDeckFile("./dumped_decks/main_deck.txt")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	dumpedDeck, readFileErr := ReadDeckFromFile("./dumped_decks/main_deck.txt")
	if readFileErr != nil {
		fmt.Println(readFileErr)
		os.Exit(1)
	}

	fmt.Println("Dumped deck: ", dumpedDeck.toString())
}
