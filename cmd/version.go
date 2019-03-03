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

	"github.com/spf13/cobra"
)

// GitSummary is populated during build time via ldflags
var GitSummary string

// Version is populated during build time via ldflags
var Version string

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"vers", "ver", "v"},
	Short:   "Output version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%v git:%v\n", Version, GitSummary)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
