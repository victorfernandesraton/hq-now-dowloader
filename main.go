package main

import (
	"github.com/victorfernandesraton/hq-now-dowloader/commands"
)

func main() {
	err := commands.CreateAllChapters(309)
	if err != nil {
		panic(err)
	}

}
