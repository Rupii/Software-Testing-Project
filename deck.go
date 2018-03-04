package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Creating a type deck

type deck []string

func newDeck() deck {
	cards := deck{}

	// named only a few cards
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValue := []string{"Ace", "two"}

	for _, suits := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suits)
		}
	}

	return cards
}

func (d deck) print() {

	for _, card := range d {
		fmt.Println(card)
	}

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(file string) error {
	return ioutil.WriteFile(file, []byte(d.toString()), 0666)
	// 0666 is just the file permission accessing all the read and write permissions
}

func deckFromFile(file string) deck {
	biteSlice, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}

	s := strings.Split(string(biteSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// we can give just the rand func but it does not shuffle in random order
	// i.e, it sets a seed value and every time it takes the same seed value to generate the rand
	// the following 2 lines are to make sure that we don't get same seed every time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	//NewSource needs a int64 value so we pass the nano sec

	for i := range d {
		position := r.Intn(len(d) - 1)

		//swap
		d[i], d[position] = d[position], d[i]
	}
}
