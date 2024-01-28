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
