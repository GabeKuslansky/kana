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
)

func Deck() *cli.Command {
	return &cli.Command{
		Name:    "deck",
		Aliases: []string{"d"},
		Usage:   "Manage decks",
		Action: func(c *cli.Context) error {
			return PickDeck()
		},
		// Subcommands: []*cli.Command{
		// 	{
		// 		Name:    "add",
		// 		Aliases: []string{"a", "new"},
		// 		Usage:   "Add a deck",
		// 		Flags: []cli.Flag{
		// 			&cli.StringFlag{
		// 				Name:    "name",
		// 				Aliases: []string{"n"},
		// 				Usage:   "Deck name",
		// 			},
		// 		},
		// 		Action: func(c *cli.Context) error {
		// 			return deck.AddFromInteractivePrompt()
		// 		},
		// 	},
		// 	{
		// 		Name:    "list",
		// 		Aliases: []string{"l"},
		// 		Usage:   "List decks",
		// 		Action: func(c *cli.Context) error {
		// 			return deck.List()
		// 		},
		// 	},
		// },
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

func AddCard() {
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

	_, err := anki.AddCard(front, back, db.Open().DefaultDeckName)
	if err != nil {
		panic(err)
	}

	if confirmChoice {
		DeckMainMenu()
	} else {
		AddCard()
	}
}
