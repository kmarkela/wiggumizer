package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kmarkela/Wiggumizeng/cmd"
	"github.com/kmarkela/Wiggumizeng/internal/parser"
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
	var bh = &parser.BrowseHistory{}
	parser.ParseHistory(data, bh)

	// check action

	// start search
	// start scan

	return

}
