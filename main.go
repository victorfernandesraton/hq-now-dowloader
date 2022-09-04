package main

import (
	"github.com/victorfernandesraton/hq-now-dowloader/commands"
	"github.com/victorfernandesraton/hq-now-dowloader/convert"
)

func main() {
	urlInfo, err := convert.ParseUrlFromHQ("https://www.hq-now.com/hq/2909/vingadores-2018")
	if err != nil {
		panic(err)
	}
	// err := commands.CreateAllChapters(309)

	err = commands.CreateAllChapters(urlInfo.ID)
	if err != nil {
		panic(err)
	}

}
