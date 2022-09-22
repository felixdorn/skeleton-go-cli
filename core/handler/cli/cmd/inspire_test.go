package cmd_test

import (
	"testing"
)

func TestNewInspireCommand(t *testing.T) {
	runTestCmd(t, []cmdTestCase{
		{
			name:   "prints a quote",
			cmd:    "inspire",
			golden: "testdata/output/prints-a-quote.golden",
		},
	})
}
