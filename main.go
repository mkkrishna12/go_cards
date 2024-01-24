package main

import "fmt"

func main() {
	// fmt.Println("Hi start this project")
	newCard := newDeck()
	newCard.saveTofile("mycard")
	newDEckFromDisk := newDeckFromFile("mycard")
	newDEckFromDisk.shuffleDeck()
	fmt.Println(newDEckFromDisk)
}
