// gitflow is a branching model for using Git within teams.
//
// Copyright 2010 Vincent Driessen.
// Copyright 2012 MediaSift Ltd.
// Copyright 2020-present Ganbaro Digital Ltd.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification are permitted provided that the following conditions are
// met:
//
//   1. Redistributions of source code must retain the above copyright
//      notice, this list of conditions and the following disclaimer.
//
//   2. Redistributions in binary form must reproduce the above copyright
//      notice, this list of conditions and the following disclaimer in the
//      documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
// IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED
// TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
// TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
// LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
// NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// The views and conclusions contained in the software and documentation
// are those of the authors and should not be interpreted as representing
// official policies, either expressed or implied, of the copyright holders
// or contributors.
//

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new git repo with support for the branching model.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	DisableFlagsInUseLine: true,
	DisableFlagParsing:    true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("init subcommand called\n\n")

		// create the args to pass into `git`
		childArgs := make([]string, len(args)+2)
		childArgs[0] = "hf"
		childArgs[1] = "init"
		for i := 0; i < len(args); i++ {
			childArgs[i+2] = args[i]
		}

		// call the shell script
		childProc := exec.Command("git", childArgs...)
		childProc.Stdin = os.Stdin
		childProc.Stdout = os.Stdout
		childProc.Stderr = os.Stderr

		err := childProc.Start()
		if err != nil {
			os.Exit(127)
		}
		err = childProc.Wait()

		// pass on the shell script's exit status code
		os.Exit(childProc.ProcessState.ExitCode())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
