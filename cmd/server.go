// Copyright © 2018 Patrick Nuckolls <nuckollsp+poorbox at gmail>
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
	//	"log"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run the poorbox api server",
	Long: `Start the poorbox api.

This program serves information about the files in poorbox
(and where to find them) to user browsers, using graphql.
It should be exposed under "http://<domain>/gq". Defaults to
localhost:8124.

Note that this program is NOT designed to serve any static
pages or any video files involved in poorbox; apache handles
those, which must be started seperatly.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
	Args: cobra.NoArgs,
}
var gqPort int

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().IntVarP(
		&gqPort,
		"server-port",
		"p",
		8124,
		"The port to run the api on")
}
