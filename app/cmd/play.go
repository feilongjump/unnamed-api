package cmd

import (
	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go playground, but running at application context",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {

}
