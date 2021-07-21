package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct {
}

type TamilBot struct {
}

func (bot EnglishBot) getGreeting() string {
	return "Welcome!"
}

func (bot TamilBot) getGreeting() string {
	return "Vanakkam!"
}

func (bot EnglishBot) getDance() int {
	return 123
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := EnglishBot{}
	printGreeting(eb)

	tb := TamilBot{}
	printGreeting(tb)
}
