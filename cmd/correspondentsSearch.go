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
	"net/url"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var correspondentsSearchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Search for a correspondent by name",
	Long: `This allows you to search for a correspondent by name.
The search uses a 'contains' search method with case sensitivity disabled by default.

Example usage:
paperless-cli correspondent search -n "hertz rental"
paperless-cli correspondent search -n dmv -s`,
	Run: func(cmd *cobra.Command, args []string) {
		name = url.QueryEscape(name)
		corrs, err := PaperInst.GetCorrespondent(name, caseSensitive)
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		fmt.Printf("%v results found:\n", len(corrs))
		for _, corr := range corrs {
			fmt.Println(corr)
		}
	},
}

func init() {
	correspondentsSearchCmd.Flags().BoolVarP(&caseSensitive, "case_sensitive", "s", false, "Enable case sensitivity")
	correspondentsSearchCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the correspondent to search for (required")
	correspondentsSearchCmd.MarkFlagRequired("name")
	correspondentsCmd.AddCommand(correspondentsSearchCmd)
}
