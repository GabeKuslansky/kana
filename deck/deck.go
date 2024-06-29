package deck

import (
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/gabekus/kana/anki"
	"github.com/gabekus/kana/db"
)

func AddFromInteractivePrompt() error {
	println("adding deck")
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
	opts := []huh.Option[string]{}
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

	num, err := strconv.Atoi(chosenDeckId)
	if err != nil {
		panic(err)
	}

	if err := db.UpdateDefaultDeck(num); err != nil {
		panic(err)
	}

	return nil
}
