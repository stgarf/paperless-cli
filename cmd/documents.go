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
)

// documentsCmd represents the documents command
var documentsCmd = &cobra.Command{
	Use:     "documents",
	Aliases: []string{"document", "docs", "doc", "d"},
	Short:   "Manage documents of Paperless instance",
	Long: `Manage documents of Paperless instance.

This includes creating, viewing, editing, and deleting documents in Paperless.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Called 'documents' with args %v", args)
		fmt.Println("documents called")
	},
}

func init() {
	rootCmd.AddCommand(documentsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// documentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// documentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
