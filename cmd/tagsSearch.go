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

var name string
var caseSensitive bool

var tagSearchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Search for tag by name",
	Args:    cobra.NoArgs,
	Long: `This allows you to search for a tag by name.
The search uses a 'contains' search method with case sensitivity disabled by default.

Example usage:
paperless-cli tag search -n taxes
paperless-cli tag search -n donation -s`,
	Run: func(cmd *cobra.Command, args []string) {
		name = url.QueryEscape(name)
		tags, err := PaperInst.GetTag(name, caseSensitive)
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		fmt.Printf("%v results found:\n", len(tags))
		for _, tag := range tags {
			fmt.Println(tag)
		}
	},
}

func init() {
	tagSearchCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the tag to search for (required)")
	tagSearchCmd.MarkFlagRequired("name")
	tagSearchCmd.Flags().BoolVarP(&caseSensitive, "case_sensitive", "s", false, "Enable case sensitivity")
	tagsCmd.AddCommand(tagSearchCmd)
}
