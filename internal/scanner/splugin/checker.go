package splugin

import "github.com/kmarkela/wiggumizer/internal/historyparser"

type Checker interface {
	GetName() string
	GetDescr() string
	GetVerbose() bool
	SetVerbose(bool) error
	Check(*historyparser.HistoryItem) (Finding, bool)
}

type Finding struct {
	Host        string
	Description string
	Evidens     string
	Details     string
}
