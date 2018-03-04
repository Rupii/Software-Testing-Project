package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 6 {
		t.Errorf("UnExpected deck size %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("unExpected card %v", d[0])
	}

	if d[len(d)-1] != "two of Hearts" {
		t.Errorf("unExpected last card %v", d[len(d)-1])
	}
}

func TestToAndFrom(t *testing.T) {
	os.Remove("_decktest")
	// _decktest is just a way of describing the file is temporary

	d := newDeck()
	d.saveToFile("_decktest")

	loadedDeck := deckFromFile("_decktest")

	if len(loadedDeck) != 6 {
		t.Errorf("unExecpted deck Size %v", len(loadedDeck))
	}

	os.Remove("_decktest")
}
