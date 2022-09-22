package cmd_test

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/mattn/go-shellwords"
	"github.com/owner/repository/core/handler/cli"
	opts "github.com/owner/repository/core/handler/cli/cli"
	"gotest.tools/v3/assert"
	"os"
	"strings"
	"testing"
)

var updateGolden = flag.Bool("update", false, "update golden files")

type cmdTestCase struct {
	name    string
	cmd     string
	golden  string
	wantErr bool
	in      *os.File
	repeat  int // 0 means 1
}

func runTestCmd(t *testing.T, tests []cmdTestCase) {
	t.Helper()

	for _, tt := range tests {
		for i := 0; i <= tt.repeat; i++ {
			t.Run(tt.name, func(t *testing.T) {
				defer resetEnv()()
				args, err := shellwords.Parse(tt.cmd)
				assert.NilError(t, err)

				buf := new(bytes.Buffer)

				root := cli.New(
					opts.WithStdout(buf),
					opts.WithStderr(buf),
					opts.When(tt.in != nil, opts.WithStdin(tt.in)),
				)

				code := root.Run(args)
				if tt.wantErr && code == 0 {
					t.Errorf("expected status code to be non-zero, got 0")

					return
				}

				if err = compare(buf.Bytes(), tt.golden); err != nil {
					t.Fatal(err)
				}
			})
		}
	}
}

func resetEnv() func() {
	origEnv := os.Environ()
	return func() {
		os.Clearenv()
		for _, pair := range origEnv {
			kv := strings.SplitN(pair, "=", 2)
			os.Setenv(kv[0], kv[1])
		}
	}
}

func compare(actual []byte, filename string) error {
	actual = bytes.Replace(actual, []byte("\r\n"), []byte("\n"), -1)
	if err := update(filename, actual); err != nil {
		return err
	}

	expected, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read testdata: %w", err)
	}

	if !bytes.Equal(actual, expected) {
		return fmt.Errorf("does not match golden file %s\n\nWANT:\n'%s'\n\nGOT:\n'%s'", filename, expected, actual)
	}

	return nil
}

func update(filename string, in []byte) error {
	if !*updateGolden {
		return nil
	}

	return os.WriteFile(filename, in, 0644)
}
