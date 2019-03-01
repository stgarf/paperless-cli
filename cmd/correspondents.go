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
	"log"

	"github.com/spf13/cobra"
)

// correspondentsCmd represents the correspondents command
var correspondentsCmd = &cobra.Command{
	Use:   "correspondents",
	Short: "Manage correspondents of Paperless instance",
	Long: `Manage correspondents of Paperless instance.

This includes adding, viewing, editing, and deleting correspondents.`,
	Run: func(cmd *cobra.Command, args []string) {
		if debugFlag {
			log.Printf("DEBUG: Called 'correspondents' with args %v\n", args)
		}
		fmt.Println("correspondents called")
	},
}

func init() {
	rootCmd.AddCommand(correspondentsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// correspondentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// correspondentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
