package cli

import (
	"github.com/gabekuslansky/kana/cmds"
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
		},
	}
}

// recalc stability after succesful review

// success = (easy / good / hard)
