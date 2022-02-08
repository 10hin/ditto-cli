package cmd

import (
	"context"
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
			log.Println("rootCmd.Run() not implemented yet")
			defer log.Println("rootCmd.Run() finished")
			// TODO: implement
			log.Panicln("Not implemented yet")
		},
	}
	cfgFile      string
	outputFormat string
)

func Execute() {
	log.Println("Execute() called")
	defer log.Println("Execute() finished")
	var err error

	err = rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}

}

func init() {
	log.Println("init() called")
	defer log.Println("init() finished")
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config/ditto-cli/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "yaml", "set output format. supported values: yaml|json ; default: yaml")

	rootCmd.AddCommand(twinCmd)
	rootCmd.AddCommand(NewGetCmd())
}

func initConfig() {
	log.Println("initConfig() called")
	defer log.Println("initConfig() finished")
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

	var cfg *config.DittoConfig
	allSettings := viper.AllSettings()
	cfg, err = config.Parse(allSettings)
	if err != nil {
		log.Fatalln(err)
	}
	context.WithValue(rootCmd.Context(), "cfg", cfg)
	var cfgBytes []byte
	cfgBytes, err = json.Marshal(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	// TODO: remove following debug log
	log.Printf("server.http.url_prefix: %s\n", (string)(cfgBytes))
}
