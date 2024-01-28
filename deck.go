package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Deck []string

const (
	SmallDeckSize = 36
	BigDeckSize   = 52
	separator     = ", "
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
	return strings.Join(d, separator)
}

func deckFromString(s string) Deck {
	return strings.Split(s, separator)
}

func (d Deck) Print() {
	fmt.Println(d.toString())
}

func (d Deck) Shuffle(r *rand.Rand) {
	r.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
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

func SaveDeckToFile(filename string, d Deck) error {
	return os.WriteFile(
		filename,
		[]byte(d.toString()),
		0666,
	)
}

func ReadDeckFromFile(filename string) (Deck, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return deckFromString(string(file)), nil
}

func RemoveDeckFile(filename string) error {
	return os.Remove(filename)
}
