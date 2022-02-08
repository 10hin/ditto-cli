package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	getThingCmd = &cobra.Command{
		Use:   "thing <thingID>",
		Short: "Get or list thing(s)",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("getThingCmd.Run() called")
			defer log.Println("getThingCmd.Run() finished")

			//var err error

			// TODO: remove following debug code
			log.Println("args:")
			for idx, arg := range args {
				log.Printf("    %2d: %s\n", idx, arg)
			}

			log.Printf("  flag --output=%s\n", outputFormat)
			log.Printf("  flag --config=%s", cfgFile)

			// TODO: remove obove debug log

			log.Println("getThingCmd.Run() not implemented yet")
			panic(fmt.Errorf("getThingCmd.Run() not implemented yet"))

		},
	}
)

func NewGetThingCmd() *cobra.Command {
	return getThingCmd
}

func getThing(cmd *cobra.Command, thingID string) {

}
