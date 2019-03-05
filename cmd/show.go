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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"s", "sh"},
	Short:   "Show the current configuration",
	Long: `Shows the current configuration file for paperless-cli.
	
The configuration displayed will change based on the value of the --config flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Called 'show' with args %v", args)
		if len(args) > 0 {
			log.Debugf("Command takes no args, ignoring: %v", args)
		}
		// If no config file found, suggest creation
		if viper.ConfigFileUsed() == "" {
			fmt.Println("No configuration file found! Try 'config create'")
		} else {
			// Display configuration information
			fmt.Println(PaperInst)
			PaperInst.ShowInstanceInformation()
		}
		log.Debug("Done calling 'show'")
	},
}

func init() {
	configCmd.AddCommand(showCmd)
}
