package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := createDeck()
	fmt.Println("error ", len(d))
	if len(d) != 521 {
		t.Errorf("Length is not as expected %v", len(d))
	}
}
