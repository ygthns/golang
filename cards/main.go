package main

func main() {
	//cards := newDeck()
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()

	cards := newDeck()
	//	cards.saveToFile("thirdTry")

	// cards := newDeckFromFile("firstTry")
	// cards.print()
	cards.shuffle()
	cards.print()

}
