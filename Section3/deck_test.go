package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expeccted deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktest"
	os.Remove(fileName)

	d := newDeck()
	d.saveToFile(fileName)

	newD := newDeckFromFile(fileName)

	if len(newD) != 16 {
		t.Errorf("Expeccted deck length of 16 but got %v", len(d))
	}

	os.Remove(fileName)
}
