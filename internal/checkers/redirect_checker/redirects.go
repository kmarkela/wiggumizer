package redirect_checker

import (
	"fmt"

	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
	"github.com/kmarkela/Wiggumizeng/internal/scanner/splugin"
)

type redirectChecker struct {
	name, descr string
	verbose     bool
}

// declare checker
func New() splugin.Checker {
	return &redirectChecker{
		name:  "redirect_checker",
		descr: "This module is searching for Redirects",
	}
}

func (rc redirectChecker) GetName() string {
	return rc.name
}

func (rc redirectChecker) GetDescr() string {
	return rc.descr
}

func (rc redirectChecker) GetVerbose() bool {
	return rc.verbose
}

func (rc *redirectChecker) SetVerbose(v bool) error {
	rc.verbose = v
	return nil
}

func (rc redirectChecker) Check(hi *historyparser.HistoryItem) (splugin.Finding, bool) {
	var f splugin.Finding

	if hi.Status < 300 || hi.Status > 399 {
		return f, false
	}

	f = splugin.Finding{
		Host:        hi.Host,
		Description: fmt.Sprintf("Redirect Found. Status: %d\n", hi.Status),
		Evidens:     fmt.Sprintf("Path: %s\n", hi.Path),
	}

	if len(hi.Req.Body) != 0 {
		f.Details = hi.Req.Body
	}

	return f, true

}
