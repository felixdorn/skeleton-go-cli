package cmd_test

import (
	"io"
	"os"
	"testing"

	"github.com/owner/repository/core/handler/cli/cli"
	"github.com/owner/repository/core/handler/cli/cmd"
	"gotest.tools/v3/assert"
)

func TestNewInspireCommand(t *testing.T) {
	pty, tty, err := os.Pipe()
	assert.NilError(t, err)

	c := cli.New(tty, tty, tty)
	inspire := cmd.NewInspireCommand(c)

	err = inspire.Execute()
	assert.NilError(t, err)

	err = tty.Close()
	assert.NilError(t, err)

	// read all from pty
	contents, err := io.ReadAll(pty)
	assert.NilError(t, err)

	assert.Equal(t, string(contents), "When there is no desire, all things are at peace. - Laozi\n")
}
