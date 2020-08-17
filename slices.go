package main
import "fmt"
func slices() { //automatically called.
	cards:= []string{"Ace of diamonds",newCard()}
	cards =append(cards,"six of spades")
	for i, card := range cards {
		fmt.Println(i,card)
	}
	
}
func newCard() string{
	return "Five of diamonds"
}
 