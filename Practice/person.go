package main

// This file used for demonstrating the pointer and variables
type contactInfo struct {
	address string
	zipCode string
}
type person struct {
	firstname string
	lastname  string
	contactInfo
}

func (personUpdate *person) updateFirstName(firstName string) {
	personUpdate.firstname = firstName
}
