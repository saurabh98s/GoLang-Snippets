package main
import ("fmt"
		"strings"	)
	//create a new type which is deck
	//this deck is slice
type deck []string

func newDeck() deck { //returns a value of type deck
 cards:= deck{}
 cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
 cardValues := []string{"Ace","Two","Three", "Four"}
  for _,suits:= range cardSuits {
  	for _,value := range cardValues{
  		cards = append(cards,suits+" of "+value)	
  	}
  	
  }
  return cards
 }

func (d deck) print() { //d is the reciever 
	for i,card:= range d {
		fmt.Println(i,card)
	}
	/* cards is type deck and type deck works
	only because the reciever is set to all function
	of type deck which is []string*/
}
func deal(d deck,handSize int) (deck ,deck) {
	return d[:handSize],d[handSize:]
 	
 }
func (d deck) toString() string {
	return strings.Join([]string(d),",") 


			
}