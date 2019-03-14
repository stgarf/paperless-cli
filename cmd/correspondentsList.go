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

var correspondentsListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Get a list of correspondents",
	Aliases: []string{"li", "l"},
	Run: func(cmd *cobra.Command, args []string) {
		corrs, err := PaperInst.GetCorrespondents()
		if err != nil {
			log.Errorf("%s", err)
		}
		fmt.Printf("%v results found:\n", len(corrs))
		for i, corr := range corrs {
			fmt.Printf("%d. %v\n", i+1, corr.Name)
		}
	},
}

func init() {
	correspondentsCmd.AddCommand(correspondentsListCmd)
}
