package secretchecker

import (
	"reflect"
	"testing"

	"github.com/kmarkela/wiggumizer/internal/historyparser"
	"github.com/kmarkela/wiggumizer/internal/scanner/splugin"
)

func TestCheck(t *testing.T) {

	tests := []struct {
		name     string
		hi       *historyparser.HistoryItem
		found    bool
		expected splugin.Finding
	}{
		{
			name:     "No found",
			hi:       &historyparser.HistoryItem{Res: historyparser.HistoryReqRes{Body: "Fastly error: unknown domain"}},
			found:    false,
			expected: splugin.Finding{},
		},
		{
			name: "AWS",
			hi: &historyparser.HistoryItem{Host: "pwn.com",
				Path: "/aws",
				Res:  historyparser.HistoryReqRes{Body: "This is the key AKIAIOSFODNN7EXAMPLE"}},

			found: true,
			expected: splugin.Finding{
				Host:        "pwn.com",
				Description: "Secrets Found",
				Evidens:     "Path: /aws",
				Details:     "Description: Identified a pattern that may indicate AWS credentials, risking unauthorized cloud resource access and data breaches on AWS platforms.\n Match: AKIAIOSFODNN7EXAMPLE\n",
			},
		},
		{
			name: "GitHub",
			hi: &historyparser.HistoryItem{Host: "pwn.com",
				Path: "/github",
				Res:  historyparser.HistoryReqRes{Body: "This is the key ghp_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}},

			found: true,
			expected: splugin.Finding{
				Host:        "pwn.com",
				Description: "Secrets Found",
				Evidens:     "Path: /github",
				Details:     "Description: Uncovered a GitHub Personal Access Token, potentially leading to unauthorized repository access and sensitive content exposure.\n Match: ghp_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n",
			},
		},
	}

	secretChecker := New()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			f, found := secretChecker.Check(test.hi)

			if found != test.found {
				t.Errorf("FAILED. Found expected: %t, returned: %t. Found: %v", test.found, found, f)

			}

			if !reflect.DeepEqual(test.expected, f) {
				t.Errorf("FAILED. expected: %v, returned: %v", test.expected, f)
			}
		})
	}

}
