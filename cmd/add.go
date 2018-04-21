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

package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add things to poorbox",
}
var dryRun bool

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().BoolVar(
		&dryRun,
		"dry-run",
		false,
		"Print out what would happen if the command was run, without making any changes.")
}
