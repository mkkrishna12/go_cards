package main

import (
	"fmt"
)

func greet(phrase string, done chan bool) {

	fmt.Println("Hello!", phrase)
	fmt.Println("--!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {

	fmt.Println("Herrrrrllo!", phrase)
	done <- true
}

func main() {
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	<-done
	<-done
	<-done
	<-done
}
