package main

func main() {
	// cards := newDec()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
	cards.saveToFile("my_cards")

}
