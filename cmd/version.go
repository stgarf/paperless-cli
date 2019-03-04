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
	"os/user"

	fqdn "github.com/Showmax/go-fqdn"
	"github.com/spf13/cobra"
)

// The following flags are populated during build time via ldflags

// BuildDate e.g.
// RFC3339 formatted UTC date	"2016-08-04T18:07:54Z"
var BuildDate string

// GitBranch e.g.
// current branch name the code is built off	"master"
var GitBranch string

// GitCommit e.g.
// short commit hash of source tree	"0b5ed7a"
var GitCommit string

// GitState e.g.
// whether there are uncommitted changes	"clean or dirty"
var GitState string

// Version e.g.
// contents of ./VERSION file, if exists, or the value passed via the -version option	"2.0.0"
var Version string

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"vers", "ver", "v"},
	Short:   "Output version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("paperless-cli v%v built on %v from git:%v-%v (%v) by %v@%v\n",
			Version,
			BuildDate,
			GitCommit,
			GitState,
			GitBranch,
			buildUser.Username,
			buildHost,
		)
	},
}

var buildHost string
var buildUser *user.User

func init() {
	buildHost = fqdn.Get()
	if buildHost == "localhost" {
		buildHost, _ = os.Hostname()
	}
	buildUser, _ = user.Current()

	rootCmd.AddCommand(versionCmd)
}
