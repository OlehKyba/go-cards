package main

import (
	"math/rand"
	"testing"
)

const (
	smallDeck = "6♣, 7♣, 8♣, 9♣, 10♣, J♣, Q♣, K♣, A♣, " +
		"6♦, 7♦, 8♦, 9♦, 10♦, J♦, Q♦, K♦, A♦, " +
		"6♥, 7♥, 8♥, 9♥, 10♥, J♥, Q♥, K♥, A♥, " +
		"6♠, 7♠, 8♠, 9♠, 10♠, J♠, Q♠, K♠, A♠"
	bigDeck = "2♣, 3♣, 4♣, 5♣, 6♣, 7♣, 8♣, 9♣, 10♣, J♣, Q♣, K♣, A♣, " +
		"2♦, 3♦, 4♦, 5♦, 6♦, 7♦, 8♦, 9♦, 10♦, J♦, Q♦, K♦, A♦, " +
		"2♥, 3♥, 4♥, 5♥, 6♥, 7♥, 8♥, 9♥, 10♥, J♥, Q♥, K♥, A♥, " +
		"2♠, 3♠, 4♠, 5♠, 6♠, 7♠, 8♠, 9♠, 10♠, J♠, Q♠, K♠, A♠"
	smallDeckAfterShuffle = "K♠, 6♠, Q♠, Q♥, 6♣, 7♦, 7♠, 9♠, 8♠, " +
		"K♥, A♦, Q♦, 8♦, K♣, 6♥, K♦, 9♣, 9♦, " +
		"7♥, J♥, 6♦, Q♣, A♥, J♣, 8♥, 10♠, A♣, " +
		"8♣, 10♣, 7♣, A♠, 10♦, J♦, 10♥, J♠, 9♥"
	filename = "./dumped_decks/_test_deck.txt"
)

func TestNewDeckGetErrorWhenPassInvalidSize(t *testing.T) {
	_, err := NewDeck(5)
	if err == nil {
		t.Errorf("Expected to get error when pass invalid size")
	}
}

func TestNewDeckWhenCreateSmallDeck(t *testing.T) {
	deck, err := NewDeck(36)
	if err != nil {
		t.Errorf("Not expected error: %v", err)
	}

	if deck.toString() != smallDeck {
		t.Errorf("Expected small deck, but got: %v", deck)
	}
}

func TestNewDeckWhenCreateBigDeck(t *testing.T) {
	deck, err := NewDeck(52)
	if err != nil {
		t.Errorf("Not expected error: %v", err)
	}

	if deck.toString() != bigDeck {
		t.Errorf("Expected small deck, but got: %v", deck)
	}
}

func TestDealWhenSizeBiggerThenDeck(t *testing.T) {
	deck, err := NewDeck(36)
	if err != nil {
		t.Errorf("Not expected error: %v", err)
	}

	_, _, err = Deal(deck, 52)
	if err == nil {
		t.Errorf("Expected to get error when pass invalid size")
	}
}

func TestDeal(t *testing.T) {
	deck, err := NewDeck(36)
	if err != nil {
		t.Errorf("Not expected error: %v", err)
	}

	hand1, hand2, dealErr := Deal(deck, 6)
	if dealErr != nil {
		t.Errorf("Not expected error: %v", dealErr)
	}

	if len(hand1) != 6 {
		t.Errorf("Expected to get hand with 6 cards, but got: %v", hand1)
	}
	if len(hand2) != 30 {
		t.Errorf("Expected to get hand with 30 cards, but got: %v", hand2)
	}
}

func TestDeck_Shuffle(t *testing.T) {
	deck, err := NewDeck(36)
	if err != nil {
		t.Errorf("Not expected error: %v", err)
	}

	r := rand.New(rand.NewSource(1))
	deck.Shuffle(r)

	if deck.toString() != smallDeckAfterShuffle {
		t.Errorf("Expected to shuffeled deck, but got: %v", deck)
	}
}

func TestSaveDeckToFileAndReadDeckFromFile(t *testing.T) {
	deck, err := NewDeck(36)
	if err != nil {
		t.Errorf("Not expected error: %v", err)
	}

	err = SaveDeckToFile(filename, deck)
	if err != nil {
		t.Errorf("Not expected error: %v", err)
	}
	defer func() {
		err = RemoveDeckFile(filename)
		if err != nil {
			t.Errorf("Not expected error: %v", err)
		}
	}()

	dumpedDeck, readFileErr := ReadDeckFromFile(filename)
	if readFileErr != nil {
		t.Errorf("Not expected error: %v", err)
	}

	if deck.toString() != dumpedDeck.toString() {
		t.Errorf("Expected equal decks, but got: deck=%v, dumpedDeck=%v", deck, dumpedDeck)
	}
}
