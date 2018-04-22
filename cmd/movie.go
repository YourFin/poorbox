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

	"github.com/yourfin/poorbox/tmdb"
	"github.com/spf13/cobra"
)

// movieCmd represents the movie command
var movieCmd = &cobra.Command{
	Use:   "movie",
	Short: "Add movies",
	Long: `Adds movies to the database given their file names,
and moves them to the appropriate location in the filesystem`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, err := parseTmdbApiSecret()
		if err != nil {
			log.Fatal(err)
		}
		tmdb.ApiInit(apiKey)
		for _, name := range args {
			movie, _ := tmdb.CmdMovieSearch(name)
			fmt.Printf("%+v\n", movie)
		}
	},
}

func init() {
	addCmd.AddCommand(movieCmd)

	// Here you will define your flags and configuration settings.
	movieCmd.Flags().StringVarP(
		&tmdbApiKey,
		"api-key",
		"a",
		"./tmdb-secret",
		`Path to the file containing a tmdb api key`)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// movieCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
