package cmd

import (
	"flag"
	"os"
	"testing"
)

func TestParams(t *testing.T) {

	tests := []struct {
		name        string
		args        []string
		expected    Params
		expectedErr bool
	}{
		{
			// Normal params
			name:     "Params OK",
			args:     []string{"-f", "example.xml", "-o", "output.md", "-a", "scan"},
			expected: Params{History: "example.xml", Output: "output.md", Action: Scan},
		},
		{
			// def action
			name:     "def action",
			args:     []string{"-f", "example.xml", "-o", "output.md"},
			expected: Params{History: "example.xml", Output: "output.md", Action: Scan},
		},
		{
			// search action
			name:     "search action",
			args:     []string{"-f", "example.xml", "-o", "output.md", "-a", "search"},
			expected: Params{History: "example.xml", Output: "output.md", Action: Search},
		},
		{
			// def action
			name:     "def hitory",
			args:     []string{"-o", "output.md"},
			expected: Params{History: "history.xml", Output: "output.md", Action: Scan},
		},
		{
			// def output
			name:     "def output",
			args:     []string{"-f", "example.xml"},
			expected: Params{History: "example.xml", Output: "report.md", Action: Scan},
		},
		{
			// def all
			name:     "no args",
			args:     []string{},
			expected: Params{History: "history.xml", Output: "report.md", Action: Scan},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			// Clear existing flag definitions
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

			os.Args = []string{"parsertest"}
			os.Args = append(os.Args, test.args...)

			p, err := newParams()
			if err != nil && !test.expectedErr {
				t.Errorf("Error occurred: %v", err)
				return
			}

			if p.Action != test.expected.Action || p.History != test.expected.History || p.Output != test.expected.Output {
				t.Errorf("unxpected params. Got %v, expected: %v", p, test.expected)
			}
		})
	}
}
