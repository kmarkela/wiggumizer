package parser

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestParseHistory(t *testing.T) {

	// read history file
	data, err := os.ReadFile("test/history.xml")
	if err != nil {
		log.Fatalf(err.Error())
	}

	// parser history
	var bh = &BrowseHistory{}
	ParseHistory(data, bh)

	fmt.Println(bh.ListOfHosts.Keys())

	for _, v := range bh.HistoryItems {
		fmt.Println(v.Res.ContentType)
	}

}
