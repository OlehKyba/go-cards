package main

import "testing"

const (
	smallDeck = "6♣, 7♣, 8♣, 9♣, 10♣, J♣, Q♣, K♣, A♣, " +
		"6♦, 7♦, 8♦, 9♦, 10♦, J♦, Q♦, K♦, A♦, " +
		"6♥, 7♥, 8♥, 9♥, 10♥, J♥, Q♥, K♥, A♥, " +
		"6♠, 7♠, 8♠, 9♠, 10♠, J♠, Q♠, K♠, A♠"
	bigDeck = "2♣, 3♣, 4♣, 5♣, 6♣, 7♣, 8♣, 9♣, 10♣, J♣, Q♣, K♣, A♣, " +
		"2♦, 3♦, 4♦, 5♦, 6♦, 7♦, 8♦, 9♦, 10♦, J♦, Q♦, K♦, A♦, " +
		"2♥, 3♥, 4♥, 5♥, 6♥, 7♥, 8♥, 9♥, 10♥, J♥, Q♥, K♥, A♥, " +
		"2♠, 3♠, 4♠, 5♠, 6♠, 7♠, 8♠, 9♠, 10♠, J♠, Q♠, K♠, A♠"
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
	deck, err := NewDeck(54)
	if err != nil {
		t.Errorf("Not expected error: %v", err)
	}

	if deck.toString() != bigDeck {
		t.Errorf("Expected small deck, but got: %v", deck)
	}
}
