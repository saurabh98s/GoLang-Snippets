package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type which is deck
//this deck is slice of String
type deck []string

func newDeck() deck { //returns a value of type deck
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suits := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suits+" of "+value)
		}

	}
	return cards
}

func (d deck) print() { //d is the reciever
	for i, card := range d {
		fmt.Println(i, card)
	}
	/* cards is type deck and type deck works
	only because the reciever is set to all function
	of type deck which is []string*/
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		//option 1= log the error and return a call to deck

		//option2= log the error and exit the program.
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") //type conversion from byte to string.
	// s is using the split to split a deck whereever there is a comma
	return deck(s)
}
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) //sets time as seed value to generate new seed each time
	  //setting seed value so that it generates values from a new sequence
	r := rand.New(source)
	for i:= range d { //swapping at random position
		newPosition := r.Intn(len(d) - 1) //random no generator
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
 