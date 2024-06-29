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
				// Reveal option to study, add, and manage cards
			} else {
				cmds.PickDeck()
			}
			return nil
		},
	}
}

// recalc stability after succesful review

// success = (easy / good / hard)
