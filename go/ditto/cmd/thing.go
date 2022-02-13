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
			log.Println("getThingCmd.Run() called")
			defer log.Println("getThingCmd.Run() finished")

			var err error

			// TODO: remove following debug code
			log.Println("args:")
			for idx, arg := range args {
				log.Printf("    %2d: %s\n", idx, arg)
			}

			log.Printf("  flag --output=%s\n", outputFormat)
			log.Printf("  flag --config=%s", cfgFile)

			var thing *model.Thing
			thing, err = http.NewClient(config.Get().Server.HTTP).Thing().Get(args[0])
			if err != nil {
				log.Printf("request failed: %#v\n", err)
				panic(err)
			}

			var s []byte
			s, err = json.MarshalIndent(thing, "", "    ")
			fmt.Printf("%s\n", s)

		},
	}
)
