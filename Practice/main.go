package main

import (
	"fmt"
)

func workingOnMaps() {
	fmt.Println("---------Working on maps---------")
	mapObject := make(map[string][]string)
	mapObject["name"] = append(mapObject["name"], "John")

	mapObject["surname"] = append(mapObject["surname"], "Mali")
	for key, value := range mapObject {
		fmt.Println(key, value)
	}
	delete(mapObject, "surname")
	fmt.Println(mapObject)

	fmt.Println("---------End maps---------")
}
func main() {
	fmt.Println("Hi start this project")
	newCard := newDeck()
	newCard.saveTofile("mycard")
	krushna := person{
		firstname: "krushna",
		lastname:  "Mali",
		contactInfo: contactInfo{
			zipCode: "413516",
			address: "Mogarga",
		},
	}
	fmt.Println(krushna)
	krushna.updateFirstName("Vikram")
	fmt.Print(krushna)
	newDEckFromDisk := newDeckFromFile("mycard")
	newDEckFromDisk.shuffleDeck()
	fmt.Println(newDEckFromDisk.multipleValues())
	workingOnMaps()
	workingOnIntefaces()
	println("Working on the interface example ")
	var sq square
	sq.side = 10
	var tri triangle
	tri.base = 3
	tri.height = 4
	fmt.Println("Triangle Area: ", getTheAreaForShape(tri))
	fmt.Println("Triangle square: ", getTheAreaForShape(sq))
	channelsPractice()
}
