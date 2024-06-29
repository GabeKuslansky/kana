package cmds

import (
	"log"

	"github.com/gabekus/kana/card"
	cli "github.com/urfave/cli/v2"
)

func Card() *cli.Command {
	return &cli.Command{
		Name:    "card",
		Aliases: []string{"c"},
		Usage:   "Add a new flashcard",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "f",
				Usage: "Front of the card",
			},
			&cli.StringFlag{
				Name:  "b",
				Usage: "Back of the card",
			},
		},
		Action: func(c *cli.Context) error {
			front := c.String("f")
			back := c.String("b")
			deck := c.String("d")

			if front != "" && back == "" {
				log.Fatal("Missing -b (back) flag")
			}
			if back != "" && front == "" {
				log.Fatal("Missing -f (front) flag")
			}

			if back != "" && front != "" {
				return card.AddFromFlags(front, back, deck)
			} else {
				return card.AddFromInteractivePrompt()
			}
		},
	}
}
