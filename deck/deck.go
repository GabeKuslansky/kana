package deck

import (
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/gabekus/kana/anki"
	"github.com/gabekus/kana/db"
)

const (
	NEW_DECK = "new"
)

func AddFromInteractivePrompt() error {
	var name string
	huh.NewInput().
		Title("What's the deck name?").
		Value(&name).
		Run()

	id, err := anki.CreateDeck(name)
	if err != nil {
		panic(err)
	}

	err = db.UpdateDefaultDeck(id)
	if err != nil {
		panic(err)
	}

	return nil
}

func AddFromFlags(name string) error {
	println("adding deck name" + name)
	return nil
}

func Pick() error {
	deckNames, err := anki.GetDeckNamesAndIds()
	if err != nil {
		panic(err)
	}
	opts := []huh.Option[string]{{Key: "(New)", Value: NEW_DECK}}
	for name, id := range deckNames {
		opts = append(opts, huh.Option[string]{Key: name, Value: strconv.Itoa(id)})
	}

	var chosenDeckId string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Pick deck").
				Options(opts...).Value(&chosenDeckId),
		),
	)
	if err = form.Run(); err != nil {
		return nil
	}

	if chosenDeckId == NEW_DECK {
		return AddFromInteractivePrompt()
	}

	id, err := strconv.Atoi(chosenDeckId)
	if err != nil {
		panic(err)
	}

	if err := db.UpdateDefaultDeck(id); err != nil {
		panic(err)
	}

	return nil
}
