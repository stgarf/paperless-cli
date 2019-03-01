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
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cfgFile == "" {
			home, _ := homedir.Dir()
			cfgFile = home + "/.paperless-cli.yaml"
		}
		// create a new configuration
		viper.Set("hostname", "localhost")
		viper.Set("use_https", false)
		viper.Set("port", 8000)
		viper.Set("root", "/api")
		if err := viper.SafeWriteConfigAs(cfgFile); err != nil {
			if os.IsNotExist(err) {
				err = viper.WriteConfigAs(cfgFile)
			} else if replace, err := cmd.Flags().GetBool("replace"); err == nil && replace {
				log.Println("Replacing existing configuration at:", cfgFile)
				err = viper.WriteConfigAs(cfgFile)
			} else {
				log.Fatalln("A configuration exists -- refusing to replace. Check flags in 'help config create'.")
			}
		}
	},
}

func init() {
	configCmd.AddCommand(createCmd)
	createCmd.Flags().BoolP("replace", "r", false, "Replace/delete an existing .paperless-cli.yaml")
}
