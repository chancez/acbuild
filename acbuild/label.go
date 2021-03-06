// Copyright 2015 The appc Authors
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

package main

import (
	"github.com/appc/acbuild/Godeps/_workspace/src/github.com/spf13/cobra"

	"github.com/appc/acbuild/lib"
)

var (
	cmdLabel = &cobra.Command{
		Use:   "label [command]",
		Short: "Manage labels",
	}
	cmdAddLabel = &cobra.Command{
		Use:     "add NAME VALUE",
		Short:   "Add a label",
		Long:    "Updates the ACI to contain a label with the given name and value. If the label already exists, its value will be changed.",
		Example: "acbuild label add arch amd64",
		Run:     runWrapper(runAddLabel),
	}
	cmdRmLabel = &cobra.Command{
		Use:     "remove NAME",
		Aliases: []string{"rm"},
		Short:   "Remove a label",
		Long:    "Updates the labels in the ACI's manifest to not include the label for the given name",
		Example: "acbuild label remove arch",
		Run:     runWrapper(runRemoveLabel),
	}
)

func init() {
	cmdAcbuild.AddCommand(cmdLabel)
	cmdLabel.AddCommand(cmdAddLabel)
	cmdLabel.AddCommand(cmdRmLabel)
}

func runAddLabel(cmd *cobra.Command, args []string) (exit int) {
	if len(args) == 0 {
		cmd.Usage()
		return 1
	}
	if len(args) != 2 {
		stderr("label add: incorrect number of arguments")
		return 1
	}

	lockfile, err := getLock()
	if err != nil {
		stderr("label add: %v", err)
		return 1
	}
	defer func() {
		if err := releaseLock(lockfile); err != nil {
			stderr("label add: %v", err)
			exit = 1
		}
	}()

	if debug {
		stderr("Adding label %q=%q", args[0], args[1])
	}

	err = lib.AddLabel(tmpacipath(), args[0], args[1])

	if err != nil {
		stderr("label add: %v", err)
		return 1
	}

	return 0
}

func runRemoveLabel(cmd *cobra.Command, args []string) (exit int) {
	if len(args) == 0 {
		cmd.Usage()
		return 1
	}
	if len(args) != 1 {
		stderr("label remove: incorrect number of arguments")
		return 1
	}

	lockfile, err := getLock()
	if err != nil {
		stderr("label remove: %v", err)
		return 1
	}
	defer func() {
		if err := releaseLock(lockfile); err != nil {
			stderr("label remove: %v", err)
			exit = 1
		}
	}()

	if debug {
		stderr("Removing label %q", args[0])
	}

	err = lib.RemoveLabel(tmpacipath(), args[0])

	if err != nil {
		stderr("label remove: %v", err)
		return 1
	}

	return 0
}
