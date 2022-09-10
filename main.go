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
			commands.GeAllChapters(info.ID)
		},
	}
	var cmdChapter = &cobra.Command{
		Use:   "ch [url from hq chapter] [number for chapter]",
		Short: "dowload hq chapter",
		Long:  "Dowload hq chapter and generate pdf with use a valid url",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			info, err := convert.ParseUrlFromHQ(args[0])
			if err != nil {
				panic(err)
			}
			commands.GetByChapter(info.ID, args[1])
		},
	}
	var rootCmd = &cobra.Command{Use: "app"}
	cmdHq.AddCommand(cmdChapter)
	rootCmd.AddCommand(cmdHq)
	rootCmd.Execute()
}
