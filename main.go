package main

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	//fmt.Println(cards.toString())
	//cards.safeToFile("my_cards")

	//cards := newDeckFromFile("my_cards")
	//cards.print()
}
