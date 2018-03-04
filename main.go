package main

func main() {
	// new deck of cards
	cards := newDeck()

	// spliting the cards to deal of hand and remaining
	hand, remain := deal(cards, 4)
	hand.print()
	remain.print()
	// saving the File to hard drive under the name hand
	remain.saveToFile("hand")

	// retriving the file from harddrive
	fromFile := deckFromFile("hand")

	fromFile.print()

	// shuffling the cards
	fromFile.shuffle()
	fromFile.print()

}
