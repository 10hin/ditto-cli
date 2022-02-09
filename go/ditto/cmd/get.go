package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"golang.10h.in/ditto/cli/pkg/ditto/client/http"
	"golang.10h.in/ditto/cli/pkg/ditto/config"
	"golang.10h.in/ditto/cli/pkg/ditto/model"
	"log"
	"os"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get or list some resource on ditto.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			log.Println("getCmd.Run() called")
			defer log.Println("getCmd.Run() finished")

			cfg, ok := (cmd.Context().Value("cfg")).(*config.DittoConfig)
			if !ok {
				panic(fmt.Errorf("failed to load config object"))
			}

			if len(args) < 1 {
				panic(fmt.Errorf("not enough argument"))
			}
			thingID := args[0]

			client := http.NewClient(cfg.Server.HTTP)
			var thing *model.Thing
			thing, err = client.Thing().GetThing(thingID)
			if err != nil {
				panic(err)
			}

			encoder := json.NewEncoder(os.Stdout)
			encoder.SetIndent("", "  ")
			err = encoder.Encode(thing)
			if err != nil {
				panic(fmt.Errorf("failed to encode response into JSON: %w", err))
			}
		},
	}
)

func NewGetCmd() *cobra.Command {
	getCmd.AddCommand(NewGetThingCmd())
	return getCmd
}
