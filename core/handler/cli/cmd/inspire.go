package cmd

import (
	"fmt"

	"github.com/owner/repository/core/handler/cli/cli"

	"github.com/spf13/cobra"
)

func runInspire(cli *cli.CLI) error {
	_, _ = fmt.Fprintln(cli.Out(), "When there is no desire, all things are at peace. - Laozi")

	return nil
}

func NewInspireCommand(cli *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspire",
		Short: "Get a random quote",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInspire(cli)
		},
	}

	return cmd
}
