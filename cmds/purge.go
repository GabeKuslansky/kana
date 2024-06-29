package cmds

import (
	"os"

	"github.com/gabekus/kana/db"
	"github.com/urfave/cli/v2"
)

func Purge() *cli.Command {
	return &cli.Command{
		Name:    "purge",
		Aliases: []string{"p"},
		Usage:   "Purges local database",
		Action: func(c *cli.Context) error {
			path, err := db.GetDatabasePath()
			if err != nil {
				panic(err)
			}

			if db.DbFileExists() {
				if err := os.Remove(path); err != nil {
					panic(err)
				}
				println("Purged database")
			} else {
				println("No database present")
			}

			return nil
		},
	}
}
