package search

import (
	"strconv"

	"github.com/kmarkela/Wiggumizeng/cmd"
)

type sConfig struct {
	Output          outputType
	CaseInsensitive bool
}

type outputType string

const (
	outEndpoint outputType = "endpoint"
	outHeaders  outputType = "headers"
	outReqOnly  outputType = "reqOnly"
	outFull     outputType = "full"
)

func (s *Searcher) toggleCase() {

	var ops = []string{"true", "false"}

	caseInsensitive := cmd.GetSelect("caseInsensitive", ops, strconv.FormatBool(s.sConf.CaseInsensitive))

	// TODO: error handling
	s.sConf.CaseInsensitive, _ = strconv.ParseBool(caseInsensitive)

}

func (s *Searcher) toggleOutput() {

	var ops = []string{"endpoint", "headers", "reqOnly", "full"}
	output := cmd.GetSelect("Output", ops, string(s.sConf.Output))

	s.sConf.Output = outputType(output)

}

func (s *Searcher) handleConfig() {

	var ops = []string{"Output", "CaseInsensitive", "Back"}

	menuOps := cmd.GetSelect("Config", ops, "Output")

	switch menuOps {
	case "Output":
		s.toggleOutput()
	case "CaseInsensitive":
		s.toggleCase()
	case "Back":
		return
	}

	s.handleConfig()
}
