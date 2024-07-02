package cmds

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/gabekus/kana/anki"
	"github.com/gabekus/kana/db"
	cli "github.com/urfave/cli/v2"
	// "log"
)

const (
	NEW_DECK_ID       = -2
	DEFAULT_DECK_NAME = "Default"
	DEFAULT_DECK_ID   = 1
	VIEW              = 0
	ADD               = 1
	STUDY             = 2
	SWITCH            = 3
	QUIT              = 4
	BACK_ID           = -1
)

func Deck() *cli.Command {
	return &cli.Command{
		Name:    "deck",
		Aliases: []string{"d"},
		Usage:   "Manage decks",
		Action: func(c *cli.Context) error {
			return PickDeck()
		},
	}
}

func Add() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Add card",
		Action: func(c *cli.Context) error {
			return AddCard()
		},
	}
}

func View() *cli.Command {
	return &cli.Command{
		Name:    "view",
		Aliases: []string{"v"},
		Usage:   "View cards",
		Action: func(c *cli.Context) error {
			return ViewCards(db.Open().DefaultDeckName)
		},
	}
}
func DeckMainMenu() error {
	title := db.Open().DefaultDeckName
	if title == "" {
		panic("deck not found")
	}

	var action int
	huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title(fmt.Sprintf("[%s]", title)).
				Options(
					huh.NewOption("View cards", VIEW),
					huh.NewOption("Add cards", ADD),
					huh.NewOption("Study", STUDY),
					huh.NewOption("Switch deck", SWITCH),
					huh.NewOption("Quit", QUIT),
				).
				Value(&action),
		),
	).Run()

	switch action {
	case VIEW:
		ViewCards(title)
	case ADD:
		AddCard()
	case STUDY:
	case SWITCH:
		PickDeck()
	case QUIT:
		os.Exit(0)
	}
	return nil
}

type DeckOption struct {
	Id   int
	Name string
}

func PickDeck() error {
	deckNames, err := anki.GetDeckNamesAndIds()
	if err != nil {
		panic(err)
	}

	newDeckOption := DeckOption{Id: NEW_DECK_ID, Name: ""}
	opts := []huh.Option[DeckOption]{{Key: "(New)", Value: newDeckOption}}
	for name, id := range deckNames {
		shouldAddOption := true
		if id == DEFAULT_DECK_ID {
			names := [1]string{DEFAULT_DECK_NAME}
			stats, err := anki.GetDeckStats(names[:])
			if err != nil {
				panic(err)
			}

			shouldAddOption = stats[strconv.Itoa(DEFAULT_DECK_ID)].Total_In_Deck != 0
		}
		if shouldAddOption {
			option := DeckOption{Id: id, Name: name}
			opts = append(opts, huh.Option[DeckOption]{Key: name, Value: option})
		}
	}

	var chosenDeck DeckOption
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[DeckOption]().
				Title("Pick deck").
				Options(opts...).Value(&chosenDeck),
		),
	)
	if err = form.Run(); err != nil {
		return nil
	}

	if chosenDeck.Id == NEW_DECK_ID {
		return CreateDeckFromInteractivePrompt()
	}

	if err := db.UpdateDefaultDeck(chosenDeck.Id, chosenDeck.Name); err != nil {
		panic(err)
	}

	DeckMainMenu()

	return nil
}

func CreateDeckFromInteractivePrompt() error {
	var name string
	huh.NewInput().
		Title("What's the deck name?").
		Value(&name).
		Run()

	id, err := anki.CreateDeck(name)
	if err != nil {
		panic(err)
	}

	err = db.UpdateDefaultDeck(id, name)
	if err != nil {
		panic(err)
	}

	return nil
}

func CreateDeckFromFlags(name string) error {
	println("adding deck name" + name)
	return nil
}

func AddCard() error {
	var front string
	var back string
	var confirmChoice bool
	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Front").
				Value(&front),
			huh.NewInput().
				Title("Back").
				Value(&back),
			huh.NewConfirm().
				Title("Save card?").
				Affirmative("Exit").
				Negative("Save").
				Value(&confirmChoice),
		),
	).Run()

	if _, err := anki.AddCard(front, back, db.Open().DefaultDeckName); err != nil {
		return err
	}

	if confirmChoice {
		return DeckMainMenu()
	}

	return AddCard()
}

func ViewCards(title string) error {
	ids, err := anki.FindCardIds()
	if err != nil {
		return err
	}

	cards, err := anki.GetCardsInfo(ids)
	if err != nil {
		return err
	}

	maxFrontLength := 0
	for _, card := range cards {
		if len(card.Fields.Front.Value) > maxFrontLength {
			maxFrontLength = len(card.Fields.Front.Value)
		}
	}

	opts := make([]huh.Option[anki.Card], len(cards)+1)
	opts[0] = huh.Option[anki.Card]{Key: (fmt.Sprintf("\x1b[94m%-*s\x1b[0m", maxFrontLength, "(Back)")), Value: anki.Card{CardID: BACK_ID}}

	for i, card := range cards {
		yellowFront := fmt.Sprintf("\x1b[93m%-*s\x1b[0m", maxFrontLength, card.Fields.Front.Value)
		opts[i+1] = huh.Option[anki.Card]{Key: fmt.Sprintf("%s\t\t%s", yellowFront, card.Fields.Back.Value), Value: card}
	}

	var chosenCard anki.Card

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[anki.Card]().
				Title(fmt.Sprintf("[%s]", title)).
				Options(opts...).
				Value(&chosenCard),
		),
	)

	form.Run()

	if chosenCard.CardID == BACK_ID {
		return DeckMainMenu()
	}

	ManageCard(chosenCard.CardID)
	return nil
}

func ManageCard(id int64) {

}
