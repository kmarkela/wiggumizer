package splugin

import "github.com/kmarkela/Wiggumizeng/internal/historyparser"

type Checker interface {
	GetName() string
	GetDescr() string
	Check(*historyparser.HistoryItem) (Finding, bool)
}

type Finding struct {
	Host        string
	Description string
	Evidens     string
	Details     string
}
