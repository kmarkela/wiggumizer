package search

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {

	tests := []struct {
		searchStr string
		expected  sParam
	}{
		{
			searchStr: "Method POST & ReqBody *admin* & ! ResContentType HTML & ResBody success",
			expected: sParam{
				method: sMatch{value: "POST", negative: false},
				req:    sReg{body: []sMatch{{value: "*admin*", negative: false}}},
				res: sReg{body: []sMatch{{value: "success", negative: false}},
					contentType: []sMatch{{value: "HTML", negative: true}}},
			},
		},
		{
			searchStr: "Method POST & ReqBody *admin* & ! ReqBody *portal* & ! ResContentType HTML & ResBody success",
			expected: sParam{
				method: sMatch{value: "POST", negative: false},
				req:    sReg{body: []sMatch{{value: "*admin*", negative: false}, {value: "*portal*", negative: true}}},
				res: sReg{body: []sMatch{{value: "success", negative: false}},
					contentType: []sMatch{{value: "HTML", negative: true}}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.searchStr, func(t *testing.T) {

			p := parseInput(test.searchStr)

			if !reflect.DeepEqual(test.expected, p) {
				t.Errorf("returned %v is not equal expected %v", p, test.expected)
			}
		})
	}
}
