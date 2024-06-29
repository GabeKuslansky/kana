package main

import (
	// "github.com/charmbracelet/bubbletea"
	"github.com/gabekuslansky/kana/cli"
	"log"
	"os"
)

var version = "dev"

func main() {
	app := cli.App(version)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
