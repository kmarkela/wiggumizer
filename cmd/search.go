package cmd

import (
	"log"

	"github.com/kmarkela/Wiggumizeng/internal/search"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "powerfull search in browse history",
	Long: `Search parameters in interactive mode:

Avaliable search fields:
Method
ReqHeader
ReqContentType
ReqBody
ResHeader
ResContentType
ResBody

Avaliable search operators:
& - AND
! - NOT

Avaliable config flags:
-i - Case insensitive search
-br - brief output (only list uniq endpoints)
-h - only headers in output
-f - full output

Search Example:

Search for requests that satisfy the following criteria:

Request method is POST
Request body contains the term "admin"
Response content type is not HTML
Response body contains the term "success"
Make search case insensitive and output only list uniq endpoints.

ReqMethod POST & ReqBody *admin* & ! ResContentType HTML & ResBody success -br -i`,
	Run: func(cmd *cobra.Command, args []string) {

		fname, err := cmd.Flags().GetString("historyFile")
		if err != nil {
			log.Fatalln(err)
		}
		workers, err := cmd.Flags().GetInt("workers")
		if err != nil {
			log.Fatalln(err)
		}

		greet()
		bh := getHistory(fname)

		s := search.Searcher{}
		s.Search(bh, workers)

	},
}

func init() {
	searchCmd.Flags().StringP("historyFile", "f", "", "path to history file")
	searchCmd.MarkFlagRequired("historyFile")
	rootCmd.AddCommand(searchCmd)
}
