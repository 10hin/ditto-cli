package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"golang.10h.in/ditto/cli/pkg/ditto/client/http"
	"golang.10h.in/ditto/cli/pkg/ditto/config"
	"golang.10h.in/ditto/cli/pkg/ditto/model"
	"log"
)

var (
	getThingCmd = &cobra.Command{
		Use:   "thing <thingID>",
		Short: "Get or list thing(s)",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("getThingCmd.Run() called")         // TODO: remove DEBUG
			defer log.Println("getThingCmd.Run() finished") // TODO: remove DEBUG

			// TODO: remove following debug code
			log.Println("args:")
			for idx, arg := range args {
				log.Printf("    %2d: %s\n", idx, arg)
			}

			log.Printf("  flag --output=%s\n", outputFormat) // TODO: remove DEBUG
			log.Printf("  flag --config=%s", cfgFile)        // TODO: remove DEBUG

			if len(args) == 1 {
				getThing(args[0])
			} else {
				listThings(args)
			}
		},
	}
)

func getThing(thingID string) {
	var err error

	var thing *model.Thing
	thing, err = http.NewClient(config.Get().Server.HTTP).Thing().Get(thingID)
	if err != nil {
		log.Printf("request failed: %#v\n", err)
		panic(err)
	}

	var s []byte
	s, err = json.MarshalIndent(thing, "", "    ")
	fmt.Printf("%s\n", s)
}

func listThings(thingIDs []string) {
	var err error

	var things []*model.Thing
	things, err = http.NewClient(config.Get().Server.HTTP).Thing().List(thingIDs)
	if err != nil {
		log.Printf("request failed: %#v\n", err)
		panic(err)
	}

	var s []byte
	s, err = json.MarshalIndent(things, "", "    ")
	fmt.Printf("%s\n", s)
}
