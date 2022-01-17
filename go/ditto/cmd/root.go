package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			log.Panicln("Not implemented yet")
		},
	}
	cfgFile string
)

func Execute() {
	log.Println(viper.GetString("hoge"))
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "c", "config file (default is $HOME/.config/ditto-cli/config.yaml)")
	rootCmd.AddCommand(twinCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		configDir, err := os.UserConfigDir()
		cobra.CheckErr(err)

		dittoCliDir := filepath.FromSlash(configDir + "/ditto-cli")
		viper.AddConfigPath(dittoCliDir)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("dittocli")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
}
