package main

type Deck []string

func NewDeck() Deck {
	return Deck{}
}

func (d Deck) Print() {

}

func (d Deck) Shuffle() {

}

func Deal(d Deck) (Deck, Deck) {
	return Deck{}, Deck{}
}

func SaveDeckToFile(d Deck) error {
	return nil
}

func ReadDeckFromFile() (Deck, error) {
	return Deck{}, nil
}
