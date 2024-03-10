package subdchecker

import (
	"reflect"
	"testing"

	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
	"github.com/kmarkela/Wiggumizeng/internal/scanner/splugin"
)

// type HistoryItem struct {
// 	Time   string
// 	Host   string
// 	Path   string
// 	Method string
// 	Status int
// 	Req    HistoryReqRes
// 	Res    HistoryReqRes
// }

// type HistoryReqRes struct {
// 	Headers     string
// 	Body        string
// 	ContentType string
// }

// Host:        hi.Host,
// Description: fmt.Sprintf("404 Message from %s found", nfm.name),
// Evidens:     fmt.Sprintf("Path: %s\n", hi.Path),

func TestCheck(t *testing.T) {

	tests := []struct {
		name     string
		hi       *historyparser.HistoryItem
		found    bool
		expected splugin.Finding
	}{
		{
			name: "Not 404",
			hi: &historyparser.HistoryItem{Status: 200,
				Res: historyparser.HistoryReqRes{Body: "Fastly error: unknown domain"}},
			found:    false,
			expected: splugin.Finding{},
		},
		{
			name: "Fastly",
			hi: &historyparser.HistoryItem{Status: 404,
				Host: "pwn.com",
				Path: "/notfound",
				Res:  historyparser.HistoryReqRes{Body: "Fastly error: unknown domain"}},

			found: true,
			expected: splugin.Finding{
				Host:        "pwn.com",
				Description: "404 Message from Fastly found",
				Evidens:     "Path: /notfound",
			},
		},
		{
			name: "AWS",
			hi: &historyparser.HistoryItem{Status: 404,
				Host: "pwn.com",
				Path: "/notfound",
				Res:  historyparser.HistoryReqRes{Body: "The specified bucket does not exist"}},

			found: true,
			expected: splugin.Finding{
				Host:        "pwn.com",
				Description: "404 Message from AWS found",
				Evidens:     "Path: /notfound",
			},
		},

		{
			name: "404 custom",
			hi: &historyparser.HistoryItem{Status: 404,
				Host: "pwn.com",
				Path: "/notfound",
				Res:  historyparser.HistoryReqRes{Body: "page not found "}},

			found:    false,
			expected: splugin.Finding{},
		},
	}

	sbdc := New()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			f, found := sbdc.Check(test.hi)

			if found != test.found {
				t.Errorf("FAILED. Found expected: %t, returned: %t. Found: %v", test.found, found, f)

			}

			if !reflect.DeepEqual(test.expected, f) {
				t.Errorf("FAILED. expected: %v, returned: %v", test.expected, f)
			}
		})
	}

}
