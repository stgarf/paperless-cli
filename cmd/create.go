// Copyright © 2019 Steve Garf <stgarf@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c", "cr"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//replace, _ := cmd.Flags().GetBool("replace")
		log.Debugf("Called 'create' with args %v, replace: %v", args, replace)
		if cfgFile == "" {
			home, _ := homedir.Dir()
			cfgFile = home + "/.paperless-cli.yaml"
		}
		// create a new configuration
		viper.Set("hostname", "localhost")
		viper.Set("use_https", false)
		viper.Set("port", 8000)
		viper.Set("root", "/api")
		log.Debugf("Checking if a configuration exists at %v", cfgFile)
		if err := viper.SafeWriteConfigAs(cfgFile); err != nil {
			if os.IsNotExist(err) {
				log.Debugf("No configuration file found at %v", cfgFile)
				fmt.Println("No configuration exists. Creating...")
				viper.WriteConfigAs(cfgFile)
				log.Debugf("Created new configuration at %v", cfgFile)
				fmt.Println("A new configuration was created at", cfgFile)
			} else if _, err2 := os.Stat(cfgFile); err2 == nil && replace {
				log.Debugf("Replacing existing configuration at %v", cfgFile)
				viper.WriteConfigAs(cfgFile)
				fmt.Println("Replaced existing configuration")
			} else if _, err2 := os.Stat(cfgFile); err2 == nil && !replace {
				log.Debug("Configuration file already exists")
				fmt.Printf("A configuration exists at %v -- refusing to replace. Check flags in 'help config create'\n", cfgFile)
				os.Exit(1)
			} else {
				if strings.Contains(err.Error(), "extension") {
					fmt.Println(err.Error())
				} else {
					log.Fatalln(err)
				}
			}
		}
		log.Debug("Done calling 'create'")
	},
}

func init() {
	configCmd.AddCommand(createCmd)
	createCmd.Flags().BoolP("replace", "r", false, "Replace/delete an existing .paperless-cli.yaml")
}
