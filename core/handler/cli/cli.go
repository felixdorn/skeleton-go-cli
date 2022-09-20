package cli

import (
	"fmt"
	"io"

	"github.com/owner/repository/core/domain/log"
	"github.com/owner/repository/core/static"
	"github.com/spf13/cobra"
	"github.com/vite-cloud/go-zoup"
)

// Stdout provides a minimal interface for writing to stdout.
type Stdout interface {
	io.Writer
	Fd() uintptr
}

// Stdin provides a minimal interface for reading stdin.
type Stdin interface {
	io.Reader
	Fd() uintptr
}

// CLI is the command line interface for Name.
type CLI struct {
	out Stdout
	in  Stdin
	err io.Writer

	commands []*cobra.Command
}

// Out returns the current output writer.
func (c *CLI) Out() Stdout {
	return c.out
}

// In returns the current input reader.
func (c *CLI) In() Stdin {
	return c.in
}

// Err returns the current error writer.
func (c *CLI) Err() io.Writer {
	return c.err
}

// Run the name CLI with the given command line arguments.
// and returns the exit code for the command.
func (c *CLI) Run(args []string) int {
	cli := &cobra.Command{
		Use:           ":bin",
		SilenceErrors: true,
		SilenceUsage:  true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Version: fmt.Sprintf("%s, build %s", static.Version, static.Commit),
	}

	cli.SetVersionTemplate(":bin version {{.id}}\n")
	cli.SetHelpCommand(&cobra.Command{
		Use:    "__help",
		Hidden: true,
	})

	cli.AddCommand(c.commands...)

	cli.SetArgs(args)
	cli.SetOut(c.Out())
	cli.SetIn(c.In())
	cli.SetErr(c.Err())

	command := "root"
	if len(args) > 0 {
		command = args[0]
	}

	if err := cli.Execute(); err != nil {
		log.Log(zoup.InfoLevel, "command ran successfully", zoup.Fields{
			"command": command,
		})

		return 0
	} else if statusErr, ok := err.(*StatusError); ok { //nolint:errorlint // A status error would only exist at the top level of the error chain.
		_, _ = fmt.Fprintf(c.Err(), "Error: %s\n", statusErr.Status)

		log.Log(zoup.ErrorLevel, "command failed", zoup.Fields{
			"command": command,
			"err":     statusErr.Status,
			"code":    statusErr.StatusCode,
		})

		return statusErr.StatusCode
	} else {
		_, _ = fmt.Fprintf(c.Err(), "Error: %s\n", err)

		log.Log(zoup.ErrorLevel, "command failed", zoup.Fields{
			"command": command,
			"err":     err,
			"code":    1,
		})

		return 1
	}
}

// Add adds the given commands to the CLI.
func (c *CLI) Add(commands ...*cobra.Command) *CLI {
	c.commands = append(c.commands, commands...)

	return c
}

// New returns a new CLI with the given standard IO.
func New(out Stdout, in Stdin, err io.Writer) *CLI {
	return &CLI{
		out: out,
		in:  in,
		err: err,
	}
}

// StatusError is an error type that contains an exit code.
// It is used to exit with a custom exit code.
type StatusError struct {
	Status     string
	StatusCode int
}

// Error implements the error interface.
func (s *StatusError) Error() string {
	return s.Status
}
