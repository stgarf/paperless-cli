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

var documentsSearchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Search for a document by name",
	Long: `This allows you to search for a document by name.
	The search uses a 'contains' search method with case sensitivity disabled by default.
	
	Example usage:
	paperless-cli doc search -n "phone bill"
	paperless-cli documents search -n donation -s.`,
	Run: func(cmd *cobra.Command, args []string) {
		name = url.QueryEscape(name)
		docs, err := PaperInst.GetDocument(name, caseSensitive)
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		fmt.Printf("%v results found:\n", len(docs))
		for i, doc := range docs {
			fmt.Printf("%d. %v - %v\n", i+1, doc.Correspondent, doc.Title)
		}
	},
}

func init() {
	documentsCmd.AddCommand(documentsSearchCmd)
	documentsSearchCmd.Flags().BoolVarP(&caseSensitive, "case_sensitive", "s", false, "Enable case sensitivity")
	documentsSearchCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the correspondent to search for (required")
	documentsSearchCmd.MarkFlagRequired("name")
}
