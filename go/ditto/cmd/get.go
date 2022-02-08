package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get or list some resource on ditto.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("getCmd.Run() called")
			defer log.Println("getCmd.Run() finished")
		},
	}
)

func NewGetCmd() *cobra.Command {
	getCmd.AddCommand(NewGetThingCmd())
	return getCmd
}
