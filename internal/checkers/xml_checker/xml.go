package xml_checker

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
	"github.com/kmarkela/Wiggumizeng/internal/scanner/splugin"
)

type xmlChecker struct {
	name, descr string
	verbose     bool
}

// declare checker
func New() splugin.Checker {
	return &xmlChecker{
		name:  "xml_checker",
		descr: "This module is searching for XML in request parameters",
	}
}

func (xc xmlChecker) GetName() string {
	return xc.name
}

func (xc xmlChecker) GetDescr() string {
	return xc.descr
}

func (xc xmlChecker) GetVerbose() bool {
	return xc.verbose
}

func (xc *xmlChecker) SetVerbose(v bool) error {
	xc.verbose = v
	return nil
}

func checkParams(hi *historyparser.HistoryItem) (splugin.Finding, bool) {
	var f splugin.Finding

	// cehck patterns
	rePatern := `\<.*\>`

	regex, err := regexp.Compile(rePatern)
	if err != nil {
		fmt.Printf("Error compiling regex pattern: %s\n", err)
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
		Description: "Posible XML tag in params",
		Evidens:     fmt.Sprintf("Path: %s\n", hi.Path),
		Details:     fmt.Sprintf("Body: %s\n", hi.Req.Body),
	}

	return f, true
}

func (xc xmlChecker) Check(hi *historyparser.HistoryItem) (splugin.Finding, bool) {
	var f splugin.Finding

	if strings.Contains(hi.Req.ContentType, "xml") {
		f = splugin.Finding{
			Host:        hi.Host,
			Description: "XML Content Type in Req",
			Evidens:     hi.Req.ContentType,
			Details:     fmt.Sprintf("Path: %s\n", hi.Path),
		}
	}

	if f.Host != "" {
		return f, true
	}

	if xc.verbose {
		return checkParams(hi)
	}

	return f, false

}
