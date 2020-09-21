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

package helpers

import (
	"os"
	"os/exec"
)

// CallShellExt calls the given original Gitflow extension shell code,
// and returns the exit status code from the shell script
func CallShellExt(cmd []string, args []string) int {
	// create the args to pass into `git`
	childArgs := make([]string, len(args)+len(cmd)+1)

	childArgs[0] = "hf"
	for _, x := range cmd {
		childArgs = append(childArgs, x)
	}

	for _, x := range args {
		childArgs = append(childArgs, x)
	}

	// build the command to call the shell script
	childProc := exec.Command("git", childArgs...)
	childProc.Stdin = os.Stdin
	childProc.Stdout = os.Stdout
	childProc.Stderr = os.Stderr

	// make the call to `git`
	err := childProc.Start()
	if err != nil {
		os.Exit(127)
	}
	err = childProc.Wait()

	// all done
	return childProc.ProcessState.ExitCode()
}
