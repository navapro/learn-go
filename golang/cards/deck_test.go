package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected Deck Length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of spades, but got %v", d[0])
	}

	if d[len(d) -1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Ace of spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	removeFile := func () {os.Remove("_decktesting")}
	removeFile()
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected Deck Length of 16, but got %v", len(loadedDeck))
	}
	removeFile()
}