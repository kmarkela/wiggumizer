package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kmarkela/Wiggumizeng/cmd"
	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
)

const version = "0.0.1-alpha"

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
	case cmd.Scan:
		log.Fatal("not implemented yet")
	case cmd.Search:
		log.Fatal("not implemented yet!")
	}
}
