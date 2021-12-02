package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println("------------")
	// remainingCards.print()

	// final := shuffle(cards)
	// final.print()
	cards := newDeckFromFile("my_cards")
	cards.print()
}
