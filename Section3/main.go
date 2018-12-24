package main

func main() {
	cards := newDeckFromFile("my_cards")

	cards.shuffel()
	cards.print()
}
