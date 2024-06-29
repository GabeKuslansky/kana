package main

import (
	// "github.com/charmbracelet/bubbletea"
	"log"
	"os"

	"github.com/gabekus/kana/cli"
)

var version = "dev"

func main() {
	app := cli.App(version)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
