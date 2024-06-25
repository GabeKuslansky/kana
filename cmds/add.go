package cmds

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"log"
)

func Add() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a", "new"},
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
		Action: func(cCtx *cli.Context) error {
			front := cCtx.String("f")
			back := cCtx.String("b")

			if front != "" && back == "" {
				log.Fatal("Missing -b (back) flag")
			}
			if back != "" && front == "" {
				log.Fatal("Missing -f (front) flag")
			}

			if back != "" && front != "" {
				println(fmt.Sprintf("Added! (%s \u2192 %s)", front, back))
			} else {
				println("Adding with CLI")
			}
			return nil
		},
	}
}
