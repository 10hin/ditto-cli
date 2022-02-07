package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.10h.in/ditto/cli/pkg/ditto/config"
	"log"
	"os"
	"path/filepath"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ditto",
		Short: "Eclipse Ditto(TM) CLI",
		Long:  "Eclipse Ditt(TM) CLI tool to Manage and Operate",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO
			var err error
			var cfg *config.DittoConfig
			allSettings := viper.AllSettings()
			cfg, err = config.Parse(allSettings)
			if err != nil {
				log.Fatalln(err)
			}
			var cfgBytes []byte
			cfgBytes, err = json.Marshal(cfg)
			if err != nil {
				log.Fatalln(err)
			}
			// TODO: remove following debug log
			log.Printf("server.http.url_prefix: %s\n", (string)(cfgBytes))
			log.Panicln("Not implemented yet")
		},
	}
	cfgFile string
)

func Execute() {
	var err error

	err = rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}

}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config/ditto-cli/config.yaml)")

	rootCmd.AddCommand(twinCmd)
}

func initConfig() {
	var err error
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")
	} else {
		var configDir string
		home, ok := os.LookupEnv("HOME")
		if ok {
			configDir = home + "/.config"
		} else {
			configDir, err = os.UserConfigDir()
			cobra.CheckErr(err)
		}

		dittoCliDir := filepath.FromSlash(configDir + "/ditto-cli")
		viper.AddConfigPath(dittoCliDir)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("dittocli")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
}
