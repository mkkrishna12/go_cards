package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected decl length of 16, but got %v", len(d))
	}
}

func TestSavetoDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveTofile("_decktesting")
	localDeck := newDeckFromFile("_decktesting")

	if len(localDeck) != 30 {
		t.Errorf("Ecpected 16 got %v", len(localDeck))
	}
}
