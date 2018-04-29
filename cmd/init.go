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
	"log"

	"github.com/spf13/cobra"
	pdb "github.com/yourfin/poorbox/poorboxdb"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the tables in the poorbox database.",
	Run: func(cmd *cobra.Command, args []string) {
		err := parseConnect()
		if err != nil {
			log.Fatal(err)
		}
		defer pdb.Close()
		err = pdb.InitTables()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Tables intiated successfully.")
	},
}

func init() {
	dbCmd.AddCommand(initCmd)
}
