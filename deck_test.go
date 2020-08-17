
package main
import "testing" 

//how to know what to test???

// if the cards has a certain value of cards.
func TestNewDeck(t *testing.T) {

	d:= newDeck()

	if len(d)!=16 {
		t.Errorf("EXPECTED LENGTH OF 16 GOT %v",len(d))
	}

	if  d[0] != "Spades of Ace" {
		t.Errorf("Spades of Ace expected got %v",d[0])

	}
	if  d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Clubs of Four expected got %v",d[len(d)-1])

	}	
}
func TestShuffle(t *testing.T) {
	
}