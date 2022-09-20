package cli

import (
	"os"

	"github.com/owner/repository/core/handler/cli/cli"
	"github.com/owner/repository/core/handler/cli/cmd"
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
