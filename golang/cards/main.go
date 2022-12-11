package main

func main() {
	cards := newDeckFromFile("cards")
	cards.print()
}

func newCard() string {
	return "Five of Diamond"
}

