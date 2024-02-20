package search

import (
	"fmt"

	"github.com/kmarkela/Wiggumizeng/cmd"
)

func (s *Searcher) menu() {

	ops := []string{"Help", "Config", "Back"}

	input := cmd.GetSelect("Choose an option: ", ops, "Help")

	switch input {
	case "Help":
		help()
	case "Config":
		s.handleConfig()
	default:
		return
	}
}

func help() {

	help := "\nRegexp Search: \n\n"
	help += "Avaliable search fields: \n"

	searchFields := []string{"Method", "ReqHeader", "ReqContentType", "ReqBody", "ResHeader", "ResContentType", "ResBody"}

	for _, name := range searchFields {
		help += "- " + name + "\n"
	}

	help += "\nAvaliable search operators: \n"

	help += "- & - AND\n"
	help += "- ! - NOT\n\n"

	help += "Search Example: \n"
	help += "ReqMethod POST & ReqBody *admin* & ! ResContentType HTML & ResBody success\n"

	fmt.Println(help)
}
