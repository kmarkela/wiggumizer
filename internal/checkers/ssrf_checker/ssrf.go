package ssrf_checker

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
	"github.com/kmarkela/Wiggumizeng/internal/scanner/splugin"
)

type ssrfChecker struct {
	name, descr string
}

// declare checker
func New() ssrfChecker {
	return ssrfChecker{
		name:  "ssrf_checker",
		descr: "This module is searching for URL in request parameters.",
	}
}

func (sc ssrfChecker) GetName() string {
	return sc.name
}

func (sc ssrfChecker) GetDescr() string {
	return sc.descr
}

func (sc ssrfChecker) Check(hi *historyparser.HistoryItem) (splugin.Finding, bool) {
	var f splugin.Finding

	rePatern := `(https?):\/\/[^\s\/$.?#].[^\s\/]*\/?`

	regex, err := regexp.Compile(rePatern)
	if err != nil {
		return f, false
	}

	var getParams string
	pParts := strings.SplitN(hi.Path, "?", 1)
	if len(pParts) == 2 {
		getParams = pParts[1]
	}

	if match := regex.MatchString(hi.Req.Body + getParams); !match {
		return f, false
	}

	f = splugin.Finding{
		Host:        hi.Host,
		Description: "http address found in Req.",
		Evidens:     fmt.Sprintf("Path: %s\n", hi.Path),
	}

	if len(hi.Req.Body) != 0 {
		f.Details = hi.Req.Body
	}

	return f, true

}
