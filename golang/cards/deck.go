package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// Which is a slice of string
type deck []string

// Create deck functions

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds","Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _,suit := range cardSuits{
		for _,value := range cardValues{
			cards = append(cards, value + " of "+ suit)
		}
	}

	return cards
}

func (d deck)  print() {
	for _, card := range d {
		fmt.Println( card)
	}
}


func deal(d deck, handSize int) (deck, deck)  {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	
	joinedString := strings.Join([]string(d), ",")
	return joinedString
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	strBs := string(bs)
	s := strings.Split(strBs, ",")
	return deck(s)
}

func (d deck) suffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1) 
		d[i], d[newPos] = d[newPos], d[i]
	}
}