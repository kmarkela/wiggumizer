package search

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kmarkela/wiggumizer/internal/historyparser"
	"github.com/kmarkela/wiggumizer/pkg/collections"
)

func listEdnpoints(bh *historyparser.BrowseHistory, r []int) []string {
	res := collections.Set{}
	for _, v := range r {
		res.Add(bh.HistoryItems[v].Host + bh.HistoryItems[v].Path)
	}
	return res.Keys()
}

func (s *Searcher) output(bh *historyparser.BrowseHistory, r []int) {

	green := color.New(color.FgGreen).Add(color.Bold)
	red := color.New(color.FgRed).Add(color.Bold)
	blue := color.New(color.FgBlue)
	white := color.New(color.FgWhite)

	if len(r) == 0 {
		fmt.Println(red.Sprintf("Nothing Found.\n"))
		return
	}

	fmt.Println(green.Sprintf("Found %d matches.\n", len(r)))

	if s.sParam.conf.output == brief {
		for i, v := range listEdnpoints(bh, r) {
			fmt.Println(white.Sprintf("%d. %s", i+1, v))
		}
		return
	}

	for i, v := range r {
		// print endpoints only
		fmt.Println(green.Sprintf("# Match %d. \n", i+1))
		fmt.Print(blue.Sprintf("## Endpoint: "))
		fmt.Println(white.Sprintf("%s", bh.HistoryItems[v].Host+bh.HistoryItems[v].Path))
		fmt.Print(blue.Sprintf("## Time: "))
		fmt.Println(white.Sprintf("%s", bh.HistoryItems[v].Time))

		if s.sParam.conf.output >= headers {
			fmt.Println(blue.Sprintf("## ReqHeaders: "))
			fmt.Println(white.Sprintf("%s", bh.HistoryItems[v].Req.Headers))
		}

		if s.sParam.conf.output > headers && len(bh.HistoryItems[v].Req.Body) > 0 {
			fmt.Println(blue.Sprintf("## ReqBody: "))
			fmt.Println(white.Sprintf("%s \n\n", bh.HistoryItems[v].Req.Body))
		}

		if s.sParam.conf.output == headers || s.sParam.conf.output == full {
			fmt.Println(blue.Sprintf("## ResHeaders: "))
			fmt.Println(white.Sprintf("%s", bh.HistoryItems[v].Res.Headers))
		}

		if s.sParam.conf.output == full && len(bh.HistoryItems[v].Res.Body) > 0 {
			fmt.Println(blue.Sprintf("## ResBody: "))
			fmt.Println(white.Sprintf("%s \n\n", bh.HistoryItems[v].Res.Body))
		}

	}

}

func help() {

	help := "\nRegexp Search: \n\n"
	help += "Available search fields: \n"

	searchFields := []string{"Method", "ReqHeader", "ReqContentType", "ReqBody", "ResHeader", "ResContentType", "ResBody"}

	for _, name := range searchFields {
		help += "- " + name + "\n"
	}

	help += "\nAvaliable search operators: \n"

	help += "- & # AND\n"
	help += "- ! # NOT\n\n"

	help += "\nAvaliable config flags: \n"

	help += "- `-i`  # Case insensitive search\n"
	help += "- `-br` # brief output (only list uniq endpoints)\n"
	help += "- `-h`  # only headers in output\n"
	help += "- `-f`  # full output\n"

	help += "Search Example: \n"
	help += "ReqMethod POST & ReqBody *admin* & ! ResContentType HTML & ResBody success -br -i\n"

	fmt.Println(help)
}
