package cli

import (
	"github.com/owner/repository/core/handler/cli/cli"
	"github.com/owner/repository/core/handler/cli/cmd"
)

// New returns a new cli.CLI instance.
func New(opts ...cli.Opt) *cli.CLI {
	name := cli.New(
		opts...,
	)

	name.Add(
		cmd.NewInspireCommand(name),
		//
	)

	return name
}
