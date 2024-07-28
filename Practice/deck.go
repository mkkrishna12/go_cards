package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}
func (decks deck) multipleValues() (string, string) {
	return decks[0], decks[14]
}
func (elements deck) Print() {
	for i, ele := range elements {
		fmt.Println(i, ele)
	}
}

func (d deck) byteSlice() []byte {
	sliceString := []string(d)
	newSring := strings.Join(sliceString, ",")
	return []byte(newSring)
}

func (d deck) saveTofile(fileName string) error {
	byteSlice := d.byteSlice()
	return ioutil.WriteFile(fileName, byteSlice, 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Err:", err)
		os.Exit(1)
	}
	byteSlices := string(byteSlice)
	byteSlicess := strings.Split(byteSlices, ",")
	return deck(byteSlicess)
}

func (d deck) shuffleDeck() {
	for index, _ := range d {
		random := rand.Intn(len(d) - 1)
		d[index], d[random] = d[random], d[index]
	}
}
