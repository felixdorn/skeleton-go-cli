package cli

import (
	"github.com/owner/repository/core/handler/cli"
	"github.com/owner/repository/core/handler/cmd"
	"os"
)

// New returns a new cli.CLI instance.
func New() *cli.CLI {
	name := cli.New(os.Stdout, os.Stdin, os.Stderr)

	name.Add(
		cmd.NewInspireCommand(name),
		//
	)

	return name
}
