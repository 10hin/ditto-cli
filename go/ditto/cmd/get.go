package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get or list some resource on ditto.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("not implemented yet")
			panic(errors.New("not implemented yet"))
		},
	}
)

func NewGetCmd() *cobra.Command {
	getCmd.AddCommand(getThingCmd)
	return getCmd
}
