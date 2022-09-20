package cli

import (
	"os"

	"github.com/owner/repository/core/handler/cli"
	"github.com/owner/repository/core/handler/cmd"
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
