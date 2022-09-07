package main

import (
	"github.com/spf13/cobra"
	"github.com/victorfernandesraton/hq-now-dowloader/commands"
	"github.com/victorfernandesraton/hq-now-dowloader/convert"
)

func main() {
	var cmdHq = &cobra.Command{
		Use:   "hq [url from hq]",
		Short: "dowload all hq chapters",
		Long:  "Dowload all hq chapters and generate pdf with use a valid url",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			info, err := convert.ParseUrlFromHQ(args[0])
			if err != nil {
				panic(err)
			}
			if err = commands.CreateAllChapters(info.ID); err != nil {
				panic(err)
			}
		},
	}

	var cmdChapter = &cobra.Command{
		Use:   "ch [url from hq chapter]",
		Short: "dowload hq chapter",
		Long:  "Dowload hq chapter and generate pdf with use a valid url",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			info, err := convert.ParseUrlFromChapter(args[0])
			if err != nil {
				panic(err)
			}
			if err = commands.CreateByChapter(info.ID); err != nil {
				panic(err)
			}
		},
	}
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdHq, cmdChapter)
	rootCmd.Execute()
}
