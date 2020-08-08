/*
Copyright © 2020 Sourab Pareek <sourab.pareek21@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nats-streaming-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	natServer string
	natsPort  int
	clusterID string
	channel   string
	clientID  string
)

func init() {
	// cobra.OnInitialize(initConfig)

	// Generating unique client id for the client

	clientID = uuid.New().String()

	// Setting persistent flags for nats connection configuration
	rootCmd.PersistentFlags().StringVarP(&natServer, "server", "s", "localhost", "Address of nats server")
	rootCmd.PersistentFlags().IntVarP(&natsPort, "port", "p", 4222, "Port of nats server")
	rootCmd.PersistentFlags().StringVarP(&clusterID, "cluster-id", "c", "test-cluster", "Cluster ID of the nats")
	rootCmd.PersistentFlags().StringVarP(&channel, "channel", "q", "", "Channel to which message will be published")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".nats-streaming-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".nats-streaming-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
