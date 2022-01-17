package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var twinCmd = &cobra.Command{
	Use:   "twin",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO implement
		log.Panicln("not implemented yet")
	},
}
