package deck

import (
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/gabekuslansky/kana/anki"
)

func AddFromInteractivePrompt() error {
	println("adding deck")
	return nil
}

func AddFromFlags(name string) error {
	println("adding deck name" + name)
	return nil
}

func List() error {
	deckNames, err := anki.GetDeckNamesAndIds()
	if err != nil {
		println(err)
	}
	opts := []huh.Option[string]{}
	for name, id := range deckNames {
		println(name, id)
		opts = append(opts, huh.Option[string]{Key: name, Value: strconv.Itoa(id)})
	}

	var chosenDeckId string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose deck").
				Options(opts...).Value(&chosenDeckId),
		),
	)
	err = form.Run()
	println(chosenDeckId)
	if err != nil {
		return err
	}

	return nil
}
