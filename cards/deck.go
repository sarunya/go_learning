package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//A custom type of slice of strings
type deck []string

//create a deck and return
func createDeck() deck {
	newDeck := deck{}
	symbol := []string{"Diamonds", "Spades", "Clubs", "Hearts"}
	rank := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, s := range symbol {
		for _, r := range rank {
			newDeck = append(newDeck, r+" of "+s)
		}
	}

	return newDeck
}

//attach a print function to deck as receiver
func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

//attach a deal function to deck as receiver
func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	dString := []string(d)
	return strings.Join(dString, "\n")
}

func (d deck) saveToFile(fileName string) error {
	//get the string of deck
	str := d.toString()
	fmt.Println(str, "saru printed")
	//write the string to file
	return _writeToFile(fileName, str)
}

func newDeckFromFile(fileName string) deck {
	deckStr := _readFromFile(fileName)

	//split string value by \n
	d := deck(strings.Split(deckStr, "\n"))

	return d
}

func (d deck) shuffle() {
	sourceObj := rand.NewSource(time.Now().UnixNano())
	randObj := rand.New(sourceObj)
	for index := range d {
		//get a random index
		randIndex := randObj.Intn(len(d))

		if randIndex != index {
			d[index], d[randIndex] = d[randIndex], d[index]
		}
	}
}

func _writeToFile(fileName string, data string) error {
	//convert to byte
	databyte := []byte(data)
	//write the data to file
	err := ioutil.WriteFile(fileName, databyte, 0666)
	//throw error
	fmt.Println(err, fileName, databyte)
	return err
}

func _readFromFile(fileName string) string {
	//read from file
	databyte, err := ioutil.ReadFile(fileName)

	//check if there is error
	if err != nil {
		os.Exit(1)
	}

	//if not error, convert databyte slice to string
	return string(databyte)
}
