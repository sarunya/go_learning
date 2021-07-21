package main

func main() {
	//executable goes here

	newDeck := createDeck()

	// fmt.Println("------Welcome to Cards Game------ ")
	// newDeck.print()
	// hand, newDeck := newDeck.deal(5)
	// hand.print()
	// hand.saveToFile("hand_file")

	// // readDeck := newDeckFromFile("hand_file")

	// fmt.Println("Cards from new deck")
	newDeck.shuffle()

	newDeck.print()
}
