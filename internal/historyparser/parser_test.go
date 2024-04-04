package historyparser

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
	ParseHistory(&data, bh)

	// fmt.Println(bh.ListOfHosts.Keys())

	// for _, v := range bh.HistoryItems {
	// 	fmt.Println(v.Req.ContentType)
	// }

	// for _, v := range bh.HistoryItems {
	// 	fmt.Println(v.Path)
	// }

	for _, v := range bh.HistoryItems {
		fmt.Println(v.Req.ContentType)
		if v.Req.ContentType == "application/x-www-form-urlencoded" {
			fmt.Println(v.Req.Body)
			fmt.Println(len(v.Req.Body))
			fmt.Println(v.Req.Parameters.Post)
		}
	}
}
