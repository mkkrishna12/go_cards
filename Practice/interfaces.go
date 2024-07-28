package main

import (
	"fmt"
	"net/http"
	"os"
)

type bot interface {
	getGreeting() string
}

type englishBot struct {
	name     string
	greeting string
}
type spanishBot struct {
	name     string
	greeting string
}

func (eb englishBot) getGreeting() string {
	return ("Hello there " + eb.name)
}

func (sb spanishBot) getGreeting() string {
	return ("Hola " + sb.name)
}
func (eb englishBot) printGreeting() {
	fmt.Println(eb.greeting)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func workingOnIntefaces() {
	var sb spanishBot
	sb.name = "greek"
	var eb englishBot
	eb.name = "John"
	printGreeting(eb)
	printGreeting(sb)
	hitHttpRquest()
}

func hitHttpRquest() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error Occured : ", err)
		os.Exit(1)
	}
	resp = resp
}
