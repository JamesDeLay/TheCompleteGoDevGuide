package main

import "fmt"

func main() {

	// := infer type on intialization
	card := newCard("Five of Diamonds")

	// no longer have to assign type once declared
	// card = "5 of Diamonds"

	// this will error out
	// card := "5 of Diamonds"

	fmt.Println(card)
	Test()
}

func newCard(name string) string {
	return name
}
