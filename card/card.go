package card

import (
	"fmt"
)

func AddFromInteractivePrompt() error {
	return nil
}

func AddFromFlags(front string, back string, deck string) error {
	println(fmt.Sprintf("Added! (%s \u2192 %s)", front, back))
	return nil
}

func List(deckName string) {
	println("Listing cards")
}
