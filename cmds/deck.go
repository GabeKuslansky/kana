package cmds

import (
	"github.com/gabekuslansky/kana/deck"
	cli "github.com/urfave/cli/v2"
	// "log"
)

func Deck() *cli.Command {
	return &cli.Command{
		Name:    "deck",
		Aliases: []string{"d"},
		Usage:   "Manage decks",
		Action: func(cCtx *cli.Context) error {
			println("test")
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a", "new"},
				Usage:   "Add a deck",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Deck name",
					},
				},
				Action: func(c *cli.Context) error {
					return deck.AddFromInteractivePrompt()
				},
			},
			{
				Name:    "List",
				Aliases: []string{"l"},
				Usage:   "List decks",
				Action: func(c *cli.Context) error {
					return deck.List()
				},
			},
		},
	}
}
