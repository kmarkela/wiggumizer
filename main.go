package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kmarkela/Wiggumizeng/cmd"
	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
	"github.com/kmarkela/Wiggumizeng/internal/scanner"
	"github.com/kmarkela/Wiggumizeng/internal/search"
)

const version = "0.1.0-alpha"

func main() {

	// init cmd
	wgr, err := cmd.NewWigomiser()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// print version and exit
	if wgr.Params.Action == cmd.Version {
		fmt.Printf("Wiggumizer: %s", version)
		return
	}

	// print ASCII art
	wgr.Greet()

	// read history file
	data, err := os.ReadFile(wgr.Params.History)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// parser history
	var bh = &historyparser.BrowseHistory{}
	historyparser.ParseHistory(&data, bh)

	// define scope
	sh := cmd.GetMUltipleChoices("Choose hosts in Scope:", bh.ListOfHosts.Keys())
	bh.FilterByHost(sh)

	// check action
	switch wgr.Params.Action {
	case cmd.Search:
		s := search.Searcher{}
		s.Search(bh, wgr.Params.Workers)
	case cmd.Scan:
		sc := scanner.NewScanner()
		sc.Scan(bh, wgr.Params.Output, wgr.Params.Workers)
	}
}
