package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Deck []string

const (
	SmallDeckSize = 36
	BigDeckSize   = 52
)

func NewDeck(size int) (Deck, error) {
	var startRank int

	switch size {
	case SmallDeckSize:
		startRank = 6
	case BigDeckSize:
		startRank = 2
	default:
		return nil, fmt.Errorf(
			"A deck of cards can only be of the following sizes: %d, %d",
			SmallDeckSize,
			BigDeckSize,
		)
	}

	suits := [...]string{"♣", "♦", "♥", "♠"}
	highRanks := [...]string{"J", "Q", "K", "A"}
	deck := make(Deck, 0, size)

	for _, suit := range suits {
		for i := startRank; i < 11; i++ {
			rank := strconv.Itoa(i)
			deck = append(deck, rank+suit)
		}

		for _, rank := range highRanks {
			deck = append(deck, rank+suit)
		}
	}

	return deck, nil
}

func (d Deck) toString() string {
	return strings.Join(d, ", ")
}

func (d Deck) Print() {
	fmt.Println(d.toString())
}

func (d Deck) Shuffle() {

}

func Deal(d Deck, size int) (Deck, Deck, error) {
	deckLen := len(d)
	if size > deckLen {
		return nil, nil, fmt.Errorf(
			"size bigger than deck: size=%d, len(deck)=%d",
			size,
			deckLen,
		)
	}

	return d[:size], d[size:], nil
}

func SaveDeckToFile(d Deck) error {
	return nil
}

func ReadDeckFromFile() (Deck, error) {
	return Deck{}, nil
}
