package search

import (
	"fmt"

	"github.com/kmarkela/Wiggumizeng/cmd"
	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
)

type Searcher struct {
	sParam sParam
	sConf  sConfig
}

type sParam struct {
	req    sReg
	res    sReg
	method sMatch
}

type sReg struct {
	header      []sMatch
	contentType []sMatch
	body        []sMatch
}

type sMatch struct {
	value    string
	negative bool
}

func (s *Searcher) Search(bh *historyparser.BrowseHistory) {

	fmt.Print("Regexp Search. Type \"menu\" to get Search menu or \"exit\" to exit \n")

	input := cmd.GetString("Type search query: ")

	switch input {
	case "menu", "Menu", "MENU":
		s.menu()
	case "exit", "Exit", "EXIT":
		return
	case "":
		// just skip empty str
	default:
		s.sParam = parseInput(input)
		// s.doSearch(input)
	}

	// run self again
	s.Search(bh)
}
