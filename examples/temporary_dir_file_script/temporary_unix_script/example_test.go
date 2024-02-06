package example

import (
	"os/exec"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_TemporaryUnixScript(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	unixExeFile := chk.CreateTmpUnixScript([]string{`
    #!/bin/bash

    echo "Running script $0"
    `,
	})

	cmd := exec.Command(unixExeFile)
	out, err := cmd.Output()
	chk.NoErr(err)
	chk.NotNil(out)
	chk.Str(string(out), "Running script "+unixExeFile+"\n")
}
