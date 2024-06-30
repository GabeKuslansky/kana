package cli

import (
	"github.com/gabekus/kana/cmds"
	"github.com/gabekus/kana/db"
	"github.com/urfave/cli/v2"
)

func App(version string) cli.App {
	return cli.App{
		Name:    "🗡  kana",
		Version: version,
		Usage:   "Terminal based flashcard system",
		Commands: []*cli.Command{
			cmds.Deck(),
			cmds.Card(),
			cmds.Purge(),
		},
		Action: func(c *cli.Context) error {
			_db := db.Open()
			if _db.DefaultDeckId != -1 {
				// Check if deck still exists
				cmds.DeckMainMenu()
			} else {
				cmds.PickDeck()
				cmds.DeckMainMenu()
			}
			return nil
		},
	}
}

// recalc stability after succesful review

// success = (easy / good / hard)
