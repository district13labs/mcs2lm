package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "mc-aws",
		Short: "POC app",
		Long: `An application to prove that we can launch an 
AWS EC2 instance with a game server running on it which
is reachable through a configurable port.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//cobra.OnInitialize(initConfig)

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is APP_DIR/config.yml)")
}

func initConfig() {
	viper.SetConfigType("yaml")
	if cfgFile != "" {
		yamlFile, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			fmt.Printf("Error parsing YAML file: %s\n", err)
			panic(err)
		}
		// Use config file from the flag.
		viper.ReadConfig(bytes.NewBuffer(yamlFile))
	} else {
		panic("No config file path given create and pass the path with --config YOUR_CONFIG_FILE_PATH")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
