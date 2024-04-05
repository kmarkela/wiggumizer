package fuzz

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/kmarkela/Wiggumizeng/pkg/collections"
)

type Fuzzer struct {
	fuzzHistory                                        fhistory
	workers, maxReq                                    int
	excludeEndpoint, excludeParam, parameter, wordlist []string
	verbose                                            bool
	headers                                            map[string]string
	tr                                                 *http.Transport
}

func New(workers, maxReq int, headers, excludeEndpoint, excludeParam, parameter []string, filename, proxy string, v bool) (*Fuzzer, error) {

	// read wordlist
	wordlist, err := pwlist(filename)
	if err != nil {
		return nil, err
	}

	// parse headers
	h, err := pheaders(headers)
	if err != nil {
		return nil, err
	}

	// parse proxy
	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		return nil, err
	}

	return &Fuzzer{
		fuzzHistory:     fhistory{h: make(map[string]collections.Set)},
		wordlist:        wordlist,
		headers:         h,
		workers:         workers,
		maxReq:          maxReq,
		excludeEndpoint: excludeEndpoint,
		excludeParam:    excludeParam,
		parameter:       parameter,
		verbose:         v,
		tr:              &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	}, nil
}

func pwlist(filename string) ([]string, error) {

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// TODO: remove EOF
	lines := strings.Split(string(content), "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("%s is empty", filename)
	}
	return lines, nil

}

func pheaders(headers []string) (map[string]string, error) {
	rh := make(map[string]string)

	for _, h := range headers {
		p := strings.Split(h, ":")
		if len(p) < 2 {
			return nil, fmt.Errorf("%s is wrong header format", h)
		}
		rh[strings.TrimSpace(p[0])] = p[1]
	}

	return rh, nil
}

// func (f *Fuzzer) Run(bh *historyparser.BrowseHistory) {

// 	for _, v := range bh.HistoryItems {
// 		v.Host + strings.Split(v.Path, "?")[0]
// 	}
// }
