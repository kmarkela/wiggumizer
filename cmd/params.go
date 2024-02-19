package cmd

import (
	"flag"
	"fmt"
	"strings"
)

type baseAction int

// avaliable actions
const (
	Scan baseAction = iota
	Search
	Version
)

type Params struct {
	History string
	Output  string
	Action  baseAction
}

func newParams() (Params, error) {
	var p Params

	if err := p.parse(); err != nil {
		return p, err
	}

	return p, nil
}

func (p *Params) parse() error {
	var actionStr string
	var ver bool

	flag.StringVar(&p.History, "f", "history.xml", "path to XML file with burp history")
	flag.StringVar(&p.Output, "o", "report.md", "path to output")
	flag.StringVar(&actionStr, "a", "scan", "Action. scan/search")
	flag.BoolVar(&ver, "v", false, "print version")
	flag.Parse()

	if ver {
		p.Action = Version
		return nil
	}

	switch strings.ToLower(actionStr) {
	case "scan":
		p.Action = Scan
	case "search":
		p.Action = Search
	default:
		return fmt.Errorf("unsupported action: %s", actionStr)
	}

	return nil

}
